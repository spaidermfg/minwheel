package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

/*
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go process1(&wg)
	go process(&wg)
	wg.Wait()
	fmt.Println("Processs Stop")
}

func process(wg *sync.WaitGroup) {
	time.Sleep(2 * time.Second)
	fmt.Println("Process Running")
	wg.Done()
}

func process1(wg *sync.WaitGroup) {
	time.Sleep(3 * time.Second)
	fmt.Println("Process Running1")
	wg.Done()
}
*/
//chmod -x /root/dasserver/server/magusdog
func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
	serverClock()
}

// 打印信息提示
func spinner(delay time.Duration) {
	for {
		for _, v := range `-\|/` {
			fmt.Printf("\r%c", v)
			time.Sleep(delay)
		}
	}
}

// 计算斐波那契数
func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

// 并发的clock服务
func serverClock() {
	listener, err := net.Listen("tcp", "localhost:8001")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()

	input := bufio.NewScanner(c)
	for input.Scan() {
		// _, err := io.WriteString(c, time.Now().Format(time.DateTime+"\n"))
		// if err != nil {
		// 	return
		// }
		// time.Sleep(time.Second * 1)
		go echo(c, input.Text(), 1*time.Second)
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}
