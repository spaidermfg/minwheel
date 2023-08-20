package test

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"testing"
	"time"
)

func TestWebSocket(t *testing.T) {
	wsAddr := "ws://localhost:8081/ws"

	parse, err := url.Parse(wsAddr)
	if err != nil {
		log.Fatal(err)
	}

	conn, _, err := websocket.DefaultDialer.Dial(parse.String(), nil)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	go func() {
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Fatal(err)
			}

			log.Println("==============", msg)
		}
	}()

	for {
		select {
		case <-interrupt:
			fmt.Println("Interrupt signal received. Closing connection...")
			err = conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Fatal(err)
			}
		case <-time.After(time.Second):
			conn.Close()
		}

		return
	}
}

func TestWebSocketServer(t *testing.T) {
	http.HandleFunc("/ws/server", webSocketServer)
	err := http.ListenAndServe(":6767", nil)
	if err != nil {
		log.Fatal(err)
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func webSocketServer(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for {
		time.Sleep(time.Second)
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}

		message := []byte(fmt.Sprintf("你好， %v", string(p)))
		pm, err := websocket.NewPreparedMessage(messageType, message)
		if err != nil {
			log.Fatal(err)
		}

		err = conn.WritePreparedMessage(pm)
		if err != nil {
			log.Fatal()
		}
		log.Println("[=================: Write Message!]")
	}
}

func TestWebSocketClient(t *testing.T) {
	u := url.URL{
		Scheme: "ws",
		Host:   "127.0.0.1:6767",
		Path:   "/ws/server",
	}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	a := []byte("李银河")
	for {
		err := conn.WriteMessage(websocket.TextMessage, a)
		if err != nil {
			log.Println(err)
		}

		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}

		a = p
		log.Println("Receive message:", string(p))
	}
}
