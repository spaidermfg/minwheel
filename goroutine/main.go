package main

import (
	"fmt"
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
