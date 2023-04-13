package main

import (
	"log"
	"time"
)

//基于select的多路复用
//没有任何case的select{}语句会一直阻塞下去

func main() {
	abort := make(chan struct{})
	log.Println("Commencing countdown")
	c := time.Tick(time.Second * 1)

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
