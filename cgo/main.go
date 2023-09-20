package main

//#include <stdio.h>
//#include <stdlib.h>
//#include "foo.h"
//char cArray[] = {'a', 'b', 'c', 'd', 'e'};
//int ciArray[] = {1, 2, 3, 4, 5, 6};
//char *fooo = "hello fooo";
//
//enum color {
//	RED,
//  BLUE,
//  YELLOW,
//};
//
//struct employee {
//	char *id;
//	int age;
//};
//
//void print(char *str) {
//	printf("%s\n", str);
//}
import "C"
import (
	"fmt"
	"unsafe"
)

// 在C内部分配的内存，GO中的GC无法感知到，使用后需要手动释放
// 在go中调用C函数，开销大约是调用Go函数的30倍
func main() {
	/// 字符串类型
	str := "use cgo"
	cs := C.CString(str)
	defer C.free(unsafe.Pointer(cs))
	C.print(cs)

	// c-char to go-string
	fmt.Println("string: ", C.GoString(C.fooo))

	/// 数组类型
	goArray := C.GoBytes(unsafe.Pointer(&C.cArray[0]), 5)
	fmt.Printf("GoArray: %c\n", goArray)

	goArray32 := cArrayToGoArray(unsafe.Pointer(&C.ciArray[0]), unsafe.Sizeof(C.ciArray[0]), 6)
	fmt.Println(goArray32)

	/// 自定义类型
	//枚举
	var r, b, g C.enum_color = C.RED, C.BLUE, C.YELLOW
	fmt.Println(r, b, g)

	//结构体
	id := C.CString("2345")
	defer C.free(unsafe.Pointer(id))

	p := C.struct_employee{
		id:  id,
		age: 99,
	}
	fmt.Printf("%#v\n", p)

	/// 获取类型大小
	fmt.Printf("int: %#v\n", C.sizeof_int)
	fmt.Printf("char: %#v\n", C.sizeof_char)
	fmt.Printf("struct: %#v\n", C.sizeof_struct_employee)

	/// 链接外部C库
	//gcc -c foo.c
	//ar rv libfoo.a foo.o
	fmt.Println(C.count)
	C.foo()
}

func cArrayToGoArray(cArray unsafe.Pointer, elemSize uintptr, len int) []int32 {
	goArray := make([]int32, 0)
	for i := 0; i < len; i++ {
		j := *(*int32)((unsafe.Pointer)(uintptr(cArray) + uintptr(i)*elemSize))
		goArray = append(goArray, j)
	}

	return goArray
}
