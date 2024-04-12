package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sync"
)

var sy sync.WaitGroup

func main() {
	loginData := map[string]string{
		"login_name": "admin",
		"password":   "e10adc3949ba59abbe56e057f20f883e",
	}

	jsonData, err := json.Marshal(loginData)
	if err != nil {
		log.Fatal(err)
	}

	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	client := &http.Client{Transport: tr}

	go login(client, jsonData)
	go login(client, jsonData)
}

func login(client *http.Client, jsonData []byte) {
	resp, err := client.Post("https://127.0.0.1:8095/api/mlinkClient/system/login/", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	all, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("error:", err)
	}
	log.Println("[body]:", string(all))
}
