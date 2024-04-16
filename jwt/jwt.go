package main

import (
	"github.com/golang-jwt/jwt/v5"
	"log"
	"time"
)

var secretKey = []byte("helloworld")

func main() {
	token := create()
	parse(token)
}

func create() string {
	token := jwt.New(jwt.SigningMethodHS256)

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

// 签名
// 校验
