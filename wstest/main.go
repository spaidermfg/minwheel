package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	// u := url.URL{Scheme: "wss", Host: "www.google.com", Path: "/"}
	// fmt.Printf("connecting to %s\n", u.String())

	// // 建立 WebSocket 连接
	// c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	// if err != nil {
	// 	log.Fatal("dial:", err)
	// }
	// defer c.Close()

	// // 发送一条简单的消息
	// message := []byte("Hello, Google!")
	// err = c.WriteMessage(websocket.TextMessage, message)
	// if err != nil {
	// 	log.Println("write:", err)
	// 	return
	// }

	// // 接收消息并打印出来
	// _, msg, err := c.ReadMessage()
	// if err != nil {
	// 	log.Println("read:", err)
	// 	return
	// }
	// log.Printf("received: %s\n", msg)

	go handleMessages()
	http.HandleFunc("/ws", handleWebSocket)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe error:", err)
	}

}

var clients = make(map[*websocket.Conn]bool)
var broadcase = make(chan []byte)
var upgrade = websocket.Upgrader{
	CheckOrigin: func(*http.Request) bool { return true },
}

// 接收消息
func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("Upgrade error:", err)
	}
	defer ws.Close()

	clients[ws] = true

	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			delete(clients, ws)
			log.Fatal("Readmessage error:", err)
		}
		log.Println("send message:", string(message))
		broadcase <- message
	}
}

// 发送消息
func handleMessages() {
	for {
		for client := range clients {
			if err := client.WriteMessage(websocket.TextMessage, []byte("hello")); err != nil {
				delete(clients, client)
				log.Fatal("WriteMessage error:", err)
			}
		}
		time.Sleep(time.Second * 2)
	}
}
