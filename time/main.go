package main

import (
	"log"
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

	one := time.NewTicker(1 * time.Second)
	two := time.NewTicker(5 * time.Second)
	three := time.NewTicker(10 * time.Second)
	four := time.NewTicker(15 * time.Second)
	five := time.NewTicker(20 * time.Second)
	six := time.NewTicker(30 * time.Second)

	for {
		select {
		case <-one.C:
			log.Println("1===========================")
		case <-two.C:
			log.Println("====2=======================")
		case <-three.C:
			log.Println("==========3=================")
		case <-four.C:
			log.Println("================4===========")
		case <-five.C:
			log.Println("=======================5====")
		case <-six.C:
			log.Println("===========================6")
		}
	}
}
