package main

import (
	"fmt"
	"log"
	"reflect"
	"sync"
	"unsafe"
)

// unsafe
// 摆脱Go语言规则带来的限制
// 由编译器实现的包，提供一些访问语言内部特性的方法
// 类型安全：一块内存数据一旦被特定的类型所解释，它就不能再被解释为其他类型
// 使用unsafe包可以实现性能更高，与底层系统交互更容易的低级代码，但同样也有风险。
// ArbitraryType 表示一个任意表达式的类型，仅用于文档使用
// 使用场景 绕过类型保护直接操作内存，对性能敏感，与os或c交互
// 反射中TypeOf和ValueOf,将任意类型变量转换为一个interface类型，再转换为reflect.emptyInterface类型

func main() {
	log.Println("")

	//Sizeof返回操作数在内存中的字节大小
	fmt.Println(unsafe.Sizeof(float64(6)))

	//Alignof返回参数的类型需要对齐的系数
	//变量地址必须可被该变量的对齐系数整除
	fmt.Println(unsafe.Alignof(bool(true)))

	var x struct {
		a bool
		b int16
		c []float32
	}

	//返回x起始地址到b的偏移量
	//unsafe.Pointer可用于表示任意类型的指针
	fmt.Println(unsafe.Offsetof(x.c))

	//pb := &x.b
	pb := (*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))
	*pb = 42
	fmt.Println(x.b)

	through()
	reflectUnsafe()
	syncUnsafe()
	typeChange()
}

type People struct {
}

func through() {
	var a uint32 = 0x123456
	fmt.Printf("0x%x\n", a)

	p := (unsafe.Pointer)(&a)

	b := (*[4]byte)(p) // 类型穿透

	b[0] = 0x23
	b[1] = 0x24
	b[2] = 0x25
	b[3] = 0x26
	fmt.Printf("0x%x\n", a)
}

type Ide struct {
	Name  string
	Speed int32
}

func reflectUnsafe() {
	ide := &Ide{
		Name:  "clion",
		Speed: 7,
	}

	wType := reflect.TypeOf(ide)
	fmt.Println(wType.Elem().Kind().String())
}

func syncUnsafe() {
	pool := sync.Pool{}
	pool.Put("b")
	pool.Put([]int32{1, 2, 3, 434, 4})
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
}

func typeChange() {
	//byte[] -> string
	b := []byte{'a', 'b', 'c', 'd', 'e'}
	bp := unsafe.Pointer(&b)

	var s string
	s = *(*string)(bp)
	fmt.Println(s)

	//string -> byte[]
	sp := unsafe.Pointer(&s)
	var b1 []byte
	b1 = *(*[]byte)(sp)
	fmt.Println(b1)
}
