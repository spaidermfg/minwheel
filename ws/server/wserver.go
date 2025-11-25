package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许跨域（视情况而定）
	},
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("upgrade error: %v", err)
	}
	defer func() {
		_ = conn.Close()
	}()

	log.Println("client connected")

	for {
		msg := fmt.Sprintf("%v %v", time.Now().Format(time.DateTime), "Have a nice day!")
		if err = conn.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
			log.Printf("write error: %v", err)
		}

		time.Sleep(1 * time.Second)
	}
}

func main() {
	http.HandleFunc("/ws", wsHandler)
	log.Println("WebSocket server running at ws://localhost:8282/ws")
	if err := http.ListenAndServe(":8282", nil); err != nil {
		log.Fatal("server error:", err)
	}
}
