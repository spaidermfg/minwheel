package main

import (
	"fmt"
	"sync"
	"time"
)

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

//chmod -x /root/dasserver/server/magusdog
