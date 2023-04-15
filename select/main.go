package main

import (
	"log"
	"os"
	"time"
)

//基于select的多路复用
//没有任何case的select{}语句会一直阻塞下去

func main() {
	abort := make(chan struct{})

	//监听终端输入
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
	log.Println("Commencing countdown")
	c := time.Tick(time.Second * 1)

	//实现倒计时，多路复用
	for countdown := 10; countdown > 0; countdown-- {
		log.Println(countdown)
		select {
		case <-c:
			log.Println("--------------------")
		case <-abort:
			log.Println("launch abort")
			return
		}
	}
	launch()
}

func launch() {
	log.Println("Commencing countdown!!!")
}
