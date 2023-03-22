package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	//tcpClient()
	chanClient()
}

// 创建一个tcp客户端
func tcpClient() {
	conn, err := net.Dial("tcp", "localhost:8001")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)
}

// 使用通道实现客户端
func chanClient() {
	conn, err := net.Dial("tcp", "localhost:8001")
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})
	go func() {
		mustCopy(os.Stdout, conn)
		log.Println("-----done-----")
		done <- struct{}{}
	}()

	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
