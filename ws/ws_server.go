package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("升级到 WebSocket 失败:", err)
		return
	}
	defer conn.Close()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("读取消息失败:", err)
			break
		}
		log.Printf("收到消息: %s", message)
		if err := conn.WriteMessage(messageType, message); err != nil {
			log.Println("发送消息失败:", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", wsHandler)
	log.Println("WebSocket 服务器启动，监听 :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal("服务器启动失败:", err)
	}
}
