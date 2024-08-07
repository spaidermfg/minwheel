package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"strings"
	"time"
)

var secretKey = []byte("helloworld")

func main() {
	token := create()
	parse(token)

	log.Println("----------------------------------------")
	user := &UserPayload{
		Name:      "mark",
		Age:       17,
		Email:     "hello@world",
		ExpTime:   time.Hour,
		StartTime: time.Now(),
	}

	// 签名
	s, err := createToken(user)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(s)

	// 校验
	j, err := parseToken(s)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(j)
}

func create() string {
	token := jwt.New(jwt.SigningMethodHS256)
	
	//jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	//	"iat": "",
	//	"exp": "",
	//})
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = "issuer"
	claims["iat"] = time.Now().Unix()                       // 当前时间的Unix时间戳
	claims["exp"] = time.Now().Add(time.Minute * 10).Unix() // 过期时间
	claims["name"] = "audience"
	claims["email"] = "subject"
	claims["nbf"] = time.Now().Unix() // 在此时间之前，该JWT是不可用的

	signedString, err := token.SignedString(secretKey)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("token:", signedString)

	return signedString
}

func parse(token string) {
	t, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		log.Fatal(err)
	}

	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		log.Println("jwt is valid")
		log.Println(claims)
	} else {
		log.Println("jwt is invalid")
	}
}

type JWT struct {
	header    string
	payload   UserPayload
	signature string
}

type UserPayload struct {
	Name      string
	Age       int8
	Email     string
	StartTime time.Time
	ExpTime   time.Duration
}

// base64编码
func encodeBase64(data string) string {
	return base64.RawURLEncoding.EncodeToString([]byte(data))
}

// base64解码
func decodeBase64(data string) ([]byte, error) {
	decodeString, err := base64.RawURLEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	return decodeString, err
}

// 创建签名
func generateSignature(key, data []byte) (string, error) {
	hash := hmac.New(sha256.New, key)
	if _, err := hash.Write(data); err != nil {
		return "", err
	}
	return encodeBase64(string(hash.Sum(nil))), nil
}

// 创建token
func createToken(payloadData any) (string, error) {
	marshal, err := json.Marshal(payloadData)
	if err != nil {
		return "", err
	}

	header := encodeBase64(`{"alg":"HS256", "typ":"JWT"}`)
	payload := encodeBase64(string(marshal))
	headerAndPayload := strings.Join([]string{header, payload}, ".")

	signature, err := generateSignature(secretKey, []byte(headerAndPayload))
	if err != nil {
		return "", err
	}

	return strings.Join([]string{headerAndPayload, signature}, "."), err
}

// 解析token
func parseToken(token string) (*JWT, error) {
	split := strings.Split(token, ".")
	if len(split) != 3 {
		return nil, errors.New("非法Token")
	}

	header := split[0]
	payload := split[1]
	signature := split[2]

	s, err := generateSignature(secretKey, []byte(strings.Join([]string{header, payload}, ".")))
	if err != nil {
		return nil, errors.New("签名生成错误" + err.Error())
	}

	// 签名校验
	if signature != s {
		return nil, errors.New("token校验失败")
	}

	decodePayload, err := decodeBase64(payload)
	if err != nil {
		return nil, errors.New("payload解码失败" + err.Error())
	}

	user := new(UserPayload)
	if err = json.Unmarshal(decodePayload, user); err != nil {
		return nil, errors.New("user反序列化失败" + err.Error())
	}

	// 有效期校验
	log.Println(user.ExpTime, user.StartTime)
	expTime := user.StartTime.Add(user.ExpTime)
	if time.Now().After(expTime) {
		return nil, errors.New("token过期，已失效")
	} else {
		log.Println("token有效")
	}

	jwt := &JWT{
		header:    header,
		payload:   *user,
		signature: signature,
	}

	return jwt, err
}

// MiddlewareFunc ---------------------------------
type MiddlewareFunc func(header http.Handler) http.Handler

func ApplyMiddleware(handler http.Handler, middlewares ...MiddlewareFunc) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func login() {
	apiHandler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Welcome to space!")
	})
	middlewareHandler := ApplyMiddleware(apiHandler, AuthMiddleware)

	http.ListenAndServe(":9000", middlewareHandler)
}
