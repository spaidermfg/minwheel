package main

import (
	"fmt"
	"time"
)

//channel是goroutine之间互相通信的通道
//通过通信来共享内存
//创建带缓存区的通道： channel := make(chan int, 10)
//发送： send := make(chan<- int)
//接收: receive := make(<-chan int)

func main() {
	c := make(chan int, 1)
	go send(c)
	go receive(c)
	time.Sleep(3 * time.Second)
	close(c)
}

func send(c chan<- int) {
	for i := 0; i <= 10; i++ {
		c <- i
		fmt.Println("send: ", i)
	}
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println("received: ", v)
	}
}
