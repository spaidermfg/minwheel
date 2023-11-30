package main

import (
	"fmt"
	"time"
)

func main() {
	// t := time.NewTimer(10 * time.Second)
	// defer t.Stop()

	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func() {
	// 	select {
	// 	case <-t.C:
	// 		fmt.Println("asdfghjkl")
	// 	}
	// }()

	// fmt.Println("hello")
	// wg.Done()

	// one := time.NewTicker(1 * time.Second)
	// two := time.NewTicker(5 * time.Second)
	// three := time.NewTicker(10 * time.Second)
	// four := time.NewTicker(15 * time.Second)
	// five := time.NewTicker(20 * time.Second)
	// six := time.NewTicker(30 * time.Second)

	// for {
	// 	select {
	// 	case <-one.C:
	// 		log.Println("1===========================")
	// 	case <-two.C:
	// 		log.Println("====2=======================")
	// 	case <-three.C:
	// 		log.Println("==========3=================")
	// 	case <-four.C:
	// 		log.Println("================4===========")
	// 	case <-five.C:
	// 		log.Println("=======================5====")
	// 	case <-six.C:
	// 		log.Println("===========================6")
	// 	}
	// }

	// t := time.NewTicker(5 * time.Second)
	// var count int = 0
	// for i := 0; i < 10000; i++ {
	// 	count += i
	// 	log.Println("count: ", count, i)
	// }

	// //阻塞主协程
	// for {
	// 	select {
	// 	case <-t.C:
	// 		log.Println("1===========================", count)
	// 	}
	// }

	//ticker := time.NewTicker(time.Second * 5)
	//defer ticker.Stop()
	//
	//// 循环执行定时任务
	//for {
	//	select {
	//	case <-ticker.C:
	//		// 定时任务的处理逻辑
	//		fmt.Println("Execute task at", time.Now())
	//	}
	//}

	d := new(datetime)
	d.getTimeNow()

	d.timer()
}

type datetime struct{}

func (d *datetime) getTimeNow() {
	now := time.Now()
	fmt.Println(now)
}

// 单次定时器
func (d *datetime) timer() {
	time.AfterFunc(1*time.Second, func() {
		fmt.Println("timer created by after function fired!")
	})

	timer := time.NewTimer(2 * time.Second)
	defer timer.Stop()
	select {
	case <-timer.C:
		fmt.Println("timer created by new timer fired!")
	}

	select {
	case <-time.After(2 * time.Second):
		fmt.Println("timer created by after fired!")
	}

	// 复用定时器
	// 已经停止的定时器、已经触发过、Timer.C中的数据已被清空
	time1 := time.NewTimer(3 * time.Second)
	if !time1.Stop() {
		<-time1.C
	}
	time1.Reset(5 * time.Second)

	select {
	case <-time1.C:
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second * 1)
			fmt.Println("It's done two", i)
		}
	}
}
