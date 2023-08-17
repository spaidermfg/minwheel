package test

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
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
