package test

import (
	"fmt"
	"net/http"
	"sync"
	"testing"
	"time"
)

// signal 系统信号
// 一种软件中断，提供一种异步的事件处理机制，用于在操作系统内核或其他应用进程通知某一应用进程发生了某件事
// 执行系统默认处理动作
// 忽略信号
// 捕捉信号并执行自定义处理动作
// kill -l 查看平台支持的系统信号列表
// 使用kill命令将系统信号发送给某应用进程： kill -s KILL 24217(pid)

// 用户层通过Notify函数捕获的信号，Go运行时通过channel将信号发给用户层
// Notify无法捕捉SIGKILL和SIGSTOP、同步信号

func TestSignalOne(t *testing.T) {
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
