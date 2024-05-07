package main

import (
	"github.com/golang-jwt/jwt/v5"
	"log"
	"time"
)

var secretKey2 = []byte("helloworld")

func main() {
	token := create2()
	parse2(token)
}

func create2() string {
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

	signedString, err := token.SignedString(secretKey2)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("token:", signedString)

	return signedString
}

func parse2(token string) {
	t, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return secretKey2, nil
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
