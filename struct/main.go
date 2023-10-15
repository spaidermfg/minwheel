package main

import (
	"fmt"
	"time"
	"unsafe"
)

func main() {
	structAddr()
	structSet()
	structGoroutine()
	structInterface()
}

// nil struct{}
// 零内存占用，地址相同，无状态
// 实现Set集合类型，用于通道通信，作为方法接收器
func structAddr() {
	var a int
	var b int32
	var c string
	var d struct{}
	var dd struct{}
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(b))
	fmt.Println(unsafe.Sizeof(c))
	fmt.Println(unsafe.Sizeof(d))

	fmt.Printf("d: %p, dd: %p, %v \n", &d, &dd, &d == &dd)
}

type Set[K comparable] map[K]struct{}

func (s Set[K]) Add(val K) {
	s[val] = struct{}{}
}

func (s Set[K]) Remove(val K) {
	delete(s, val)
}

func (s Set[K]) Contain(val K) bool {
	_, ok := s[val]
	return ok
}

func structSet() {
	set := Set[string]{}
	set.Add("mark")
	set.Add("load")
	fmt.Println(set.Contain("mark"))
	set.Remove("mark")
	fmt.Println(set.Contain("mark"))
}

// 用于信号传递，不关心通道中传递的具体数据
func structGoroutine() {
	quit := make(chan struct{})
	go func() {
		fmt.Println("start working...")
		time.Sleep(4 * time.Second)
		close(quit)
	}()

	<-quit
	fmt.Println("finish work...")
}

// 方法接收器，不需要存储数据

type Person interface {
	Move()
	Up()
}

type SMC struct{}

func (s SMC) Move() {
	fmt.Println("move left")
}

func (s SMC) Up() {
	fmt.Println("up one")
}

func structInterface() {
	s := new(SMC)

	s.Up()
	s.Move()
}
