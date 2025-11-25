package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

var (
	serverURL string // WebSocket 服务器地址
)

func init() {
	// 解析命令行参数
	flag.StringVar(&serverURL, "url", "ws://127.0.0.1:8081/ws", "WebSocket 服务器地址")
	flag.Parse()
}

func main() {
	// 解析服务器地址
	u, err := url.Parse(serverURL)
	if err != nil {
		log.Fatalf("解析服务器地址失败: %v", err)
	}

	// 连接到 WebSocket 服务器
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatalf("连接 WebSocket 服务器失败: %v", err)
	}
	defer conn.Close()

	log.Printf("已连接到 WebSocket 服务器: %s", u.String())

	// 监听中断信号（Ctrl+C）
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// 启动一个 goroutine 读取服务器消息
	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Printf("读取消息失败: %v", err)
				return
			}
			log.Printf("收到消息: %s", message)
		}
	}()

	// 从标准输入读取用户输入并发送到服务器
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("输入消息 (或按 Ctrl+C 退出): ")
		if !scanner.Scan() {
			break
		}
		text := scanner.Text()

		// 发送消息到服务器

		err := conn.WriteMessage(websocket.TextMessage, []byte(text))
		if err != nil {
			log.Printf("发送消息失败: %v", err)
			break
		}
		log.Printf("已发送消息: %s", text)

		// 等待一段时间
		time.Sleep(100 * time.Millisecond)
	}

	// 关闭连接
	log.Println("关闭连接...")
	err = conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		log.Printf("发送关闭消息失败: %v", err)
	}
	select {
	case <-done:
	case <-time.After(time.Second):
	}
}
