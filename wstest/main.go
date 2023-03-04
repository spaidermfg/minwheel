package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

func main() {
	u := url.URL{Scheme: "wss", Host: "www.google.com", Path: "/"}
	fmt.Printf("connecting to %s\n", u.String())

	// 建立 WebSocket 连接
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	// 发送一条简单的消息
	message := []byte("Hello, Google!")
	err = c.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		log.Println("write:", err)
		return
	}

	// 接收消息并打印出来
	_, msg, err := c.ReadMessage()
	if err != nil {
		log.Println("read:", err)
		return
	}
	log.Printf("received: %s\n", msg)
}