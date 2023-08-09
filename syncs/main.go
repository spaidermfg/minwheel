package main

import (
	"log"
	"sync"
	"time"
)

// sync包：基于共享内存并发模型，互斥锁，读写锁，条件变量
// 低级同步原语

func main() {
	loopWorker()
}

// # 条件变量
// sync.Cond 一个条件变量可以理解为一个容器，这个容器中存放着一个或一组等待着某个条件成立的goroutine
// 当条件成立时，处于等待中的goroutine将会被通知运行。不用程序去轮询检查是否满足运行状态。
type signal struct{}

var ready bool

func worker(i int) {
	log.Printf("worker %d: is working...\n", i)
	time.Sleep(1 * time.Second)
	log.Printf("worker %d: works done\n", i)
}

func spawnGroup(f func(i int), num int, mu *sync.Mutex) <-chan signal {
	c := make(chan signal)
	var wg sync.WaitGroup

	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(i int) {
			for {
				mu.Lock()
				if !ready {
					mu.Unlock()
					time.Sleep(10 * time.Second)
					continue
				}
				mu.Unlock()
				log.Printf("worked %d: start to work...\n", i)
				f(i)
				wg.Done()
				return
			}
		}(i + 1)
	}

	go func() {
		wg.Wait()
		c <- signal(struct{}{})
	}()
	return c
}

func loopWorker() {
	log.Printf("start a group of workers...")
	mu := &sync.Mutex{}
	group := spawnGroup(worker, 5, mu)
	time.Sleep(5 * time.Second)
	log.Println("the group of workers start to work...")

	mu.Lock()
	ready = true
	log.Println("--------------------------------", ready)
	mu.Unlock()

	<-group
	log.Printf("The group of workers work done.")
}
