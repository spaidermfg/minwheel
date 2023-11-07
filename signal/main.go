package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type signaler struct {
}

func main() {
	s := new(signaler)
	//s.web()
	//s.notify()
	//s.double()
	//s.one()
	s.total()
}

func (s *signaler) web() {
	var wg sync.WaitGroup
	ch := make(chan error, 1)
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello, Mojoto!")
	})

	wg.Add(1)
	go func() {
		defer wg.Done()
		ch <- http.ListenAndServe("127.0.0.1:8080", nil)
	}()

	select {
	case <-time.After(2 * time.Second):
		fmt.Printf("web server started!")
	case err := <-ch:
		fmt.Println("web server start failed:", err)
	}

	wg.Wait()
	fmt.Println("web server closed!")
}

func catchAsyncSignal(c chan os.Signal) {
	for {
		s := <-c
		fmt.Println("接收到异步信号：", s)
	}
}

func triggerSyncSignal() {
	time.Sleep(3 * time.Second)
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("PANIC recover:", e)
			return
		}
	}()

	a, b := 3, 4
	fmt.Println(a * b)
}

// 用户层通过Notify函数捕获的信号，Go运行时通过channel将信号发给用户层
// Notify无法捕捉SIGKILL和SIGSTOP、同步信号
func (s *signaler) notify() {
	var wg sync.WaitGroup
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGFPE, syscall.SIGINT, syscall.SIGKILL)

	wg.Add(2)
	go func() {
		defer wg.Done()
		catchAsyncSignal(ch)
	}()

	go func() {
		defer wg.Done()
		triggerSyncSignal()
	}()

	wg.Wait()
}

// 应用进程收到异步信号后，会给每个channel发送一份异步信号副本
func (s *signaler) double() {
	ch1 := make(chan os.Signal, 1)
	ch2 := make(chan os.Signal, 1)

	signal.Notify(ch1, syscall.SIGINT, syscall.SIGTERM)
	signal.Notify(ch2, syscall.SIGINT, syscall.SIGTERM)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		a := <-ch1
		fmt.Println("c1接收到异步信号:", a)
	}()

	a := <-ch2
	fmt.Println("c2接收到异步信号:", a)
	wg.Done()

	wg.Wait()
}

// 同一个channel上两次拦截同一异步信号，会收到一个信号
func (s *signaler) one() {
	var wg sync.WaitGroup
	ch := make(chan os.Signal, 2)

	signal.Notify(ch, syscall.SIGINT)
	signal.Notify(ch, syscall.SIGINT)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			a := <-ch
			fmt.Println("ch收到异步信号：", a)
		}
	}()
	wg.Wait()
}

// 未及时处理异步信号，已有的信号会写入channel中
func (s *signaler) total() {
	ch := make(chan os.Signal, 5)
	signal.Notify(ch, syscall.SIGINT)

	time.Sleep(time.Second * 5)

	for {
		select {
		case a := <-ch:
			fmt.Println("c获取异步信号：", a)
		default:
			fmt.Println("无信号退出")
			os.Exit(0)
		}
	}
}
