package main

import (
	"fmt"
	"log"
	"unsafe"
)

// unsafe
// 摆脱Go语言规则带来的限制
// 由编译器实现的包，提供一些访问语言内部特性的方法
func main() {
	log.Println("")

	//Sizeof返回操作数在内存中的字节大小
	fmt.Println(unsafe.Sizeof(float64(6)))

	//Alignof返回参数的类型需要对齐的倍数
	fmt.Println(unsafe.Alignof(bool(true)))

	var x struct {
		a bool
		b int16
		c []float32
	}

	//返回x起始地址到b的偏移量
	fmt.Println(unsafe.Offsetof(x.c))

	//pb := &x.b
	pb := (*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))
	*pb = 42
	fmt.Println(x.b)
}

type People struct {
}
