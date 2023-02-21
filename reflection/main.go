package main

import (
	"fmt"
	"minwheel/reflection/display"
	"reflect"
	"strconv"
	"time"
)

// ----反射
// 在编译时不知道类型的情况下，可更新变量、在运行时查看值，调用方法以及直接对它们的布局进行操作。
// 何时需要：无法透视一个未知类型的布局时，需要反射的帮助。
// reflect包提供两个类型：Type和Value
func main() {
	reflect_type()
	reflect_value()
	use_any()
}

// Type是一个有很多方法的接口
// reflect.Type接口只有一个实现，reflect.TypeOf函数接收任何的interface{}参数，返回接口值对应的动态类型(reflect.Type)
// 返回的总是一个具体类型，而不是接口类型
func reflect_type() {
	t := reflect.TypeOf("dfajksdf")
	fmt.Println(t, t.String())
	fmt.Printf("%T\n", 7) //输出接口值的动态类型
	fmt.Println("--------------------")
}

// Value是一个有很多方法的接口
// reflect.ValueOf函数接收任意类型的参数，返回接口值对应的动态值(reflect.Value)
// 返回的总是一个具体值，但也可以包含接口类型
func reflect_value() {
	v := reflect.ValueOf(9)
	fmt.Println(v, v.String(), v.Int())
	fmt.Printf("%v\n", v) //输出接口值的动态值

	//调用Value的Type方法，会把它的类型以reflect.Type返回
	t := v.Type()
	fmt.Println(t.String())
	fmt.Println("--------------------")
}

// 使用类型断言来区分值的类型
func Sprint(x interface{}) string {
	type stringer interface {
		String() string
	}

	switch x := x.(type) {
	case stringer:
		return x.String()
	case string:
		return x
	case int:
		return strconv.Itoa(x)
	case bool:
		if x {
			return "true"
		}
		return "false"
	default:
		return "???"
	}
}

func use_any() {
	var x int64 = 1
	var d time.Duration = 1 * time.Nanosecond
	c := make(chan int)
	fmt.Println(display.Any(x))
	fmt.Println(display.Any([]int64{x}))
	fmt.Println(display.Any(c))
	fmt.Println(display.Any([]time.Duration{d}))
}

// Any 把任何值格式转化为string
// func Any(value interface{}) string {
// 	return formatAtom(reflect.ValueOf(value))
// }

// // formaAtom 格式化一个值，且不分析其内部结构
// func formatAtom(v reflect.Value) string {
// 	switch v.Kind() {
// 	case reflect.Invalid:
// 		return "invalid"
// 	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
// 		return strconv.FormatInt(v.Int(), 10)
// 	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
// 		return strconv.FormatUint(v.Uint(), 10)
// 	case reflect.Float32, reflect.Float64:
// 		return strconv.FormatFloat(v.Float(), 0, 0, 0)
// 	case reflect.Bool:
// 		return strconv.FormatBool(v.Bool())
// 	case reflect.String:
// 		return strconv.Quote(v.String())
// 	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Map, reflect.Slice:
// 		return v.Type().String() + " 0x" + strconv.FormatUint(uint64(v.Pointer()), 16)
// 	default:
// 		return v.Type().String() + " value"
// 	}
// }
