package chant

import (
	"fmt"
	"time"
)

func Chant() {
	nil_chan()
}

// 对没有初始化的通道进行读写操作将会发生阻塞
// fatal error: all goroutines are asleep - deadlock!
// goroutine 1 [chan receive (nil chan)]:
// 从一个已关闭的通道接收数据永远不会阻塞
func nil_chan() {
	c1 := make(chan int)
	c2 := make(chan int)
	go func() {
		time.Sleep(time.Second * 5)
		c1 <- 5
		close(c1)
	}()

	go func() {
		time.Sleep(time.Second * 7)
		c2 <- 7
		close(c2)
	}()

	for {
		select {
		case x, ok := <-c1:
			if !ok {
				c1 = nil
			} else {
				fmt.Println(x)
			}
		case x, ok := <-c2:
			if !ok {
				c2 = nil
			} else {
				fmt.Println(x)
			}
		}

		if c1 == nil && c2 == nil {
			fmt.Println("good job")
			break
		}
	}

	fmt.Println("well done")
}
