package main

import (
	"fmt"
	"reflect"
)

// ----泛型
// 使得开发者能够编写可重用、类型安全且高度通用的代码，无需提前确定代码中使用的具体类型。
// 这种类型参数化机制使得编写可重用代码更加容易，同时也提高了代码的可读性和可维护性。
// 可以使用类型参数定义通用的容器类型、函数或接口，以便在不同场景下使用。
// -支持：1.18版本推出了泛型的支持
// -优点：泛型还可以提高代码的可读性。可以编写可重用性更强的代码，不需要为每种数据类型编写独立的代码。
// -缺点：使用泛型的缺点在于可能会影响代码的性能。代码复杂难以理解。
// -类型约束：是指对泛型类型参数进行限制，以确保类型参数满足特定的条件。
// 目前支持的类型约束包括any、comparable、numeric、ordered和slice。
// "any": 表示任何类型都可以作为类型参数。使用any类型约束可以使泛型函数或类型更加通用，但是由于没有任何限制，因此也可能会影响代码的安全性。
// "comparable": 表示类型参数必须是可比较的。使用comparable类型约束可以确保类型参数支持比较运算符（如<、>、==等），并且可以保证泛型代码在运行时不会发生比较类型错误的情况。
// "numeric": 表示类型参数必须是数字类型。使用numeric类型约束可以确保类型参数支持数字运算符（如+、-、*、/等），并且可以保证泛型代码在运行时不会发生数字类型错误的情况。
// "ordered": 表示类型参数必须是可排序的。使用ordered类型约束可以确保类型参数支持排序运算符（如<、>、==等），并且可以保证泛型代码在运行时不会发生排序类型错误的情况。
// "slice": 表示类型参数必须是切片类型。使用slice类型约束可以确保类型参数是切片类型，并且可以保证泛型代码在运行时不会发生切片类型错误的情况。
// 可以使用type关键字定义一个约束类型，然后在泛型函数或类型中使用该约束类型作为类型参数的限制条件
func main() {
	use_generic()
	use_stack()
}

func use_generic() {
	ints := []int{2, 3, 2, 7, 5, 6, 1}
	i := FindMax(ints)
	fmt.Println(i)
}

func FindMax[T comparable](slice []T) T {
	max := slice[0]
	for _, v := range slice[1:] {
		if reflect.ValueOf(v).Int() > reflect.ValueOf(max).Int() {
			max = v
		}
	}
	return max
}

type Stack[T any] struct {
	items []T
}

// push stack
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// pop stack
func (s *Stack[T]) Pop() T {
	if len(s.items) == 0 {
		panic("empty stack")
	}

	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func use_stack() {
	intStack := Stack[int]{}
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)
	fmt.Println(intStack.Pop())

	strStack := Stack[string]{}
	strStack.Push("a")
	strStack.Push("b")
	strStack.Push("c")
	fmt.Println(strStack.Pop())
}

// any 任何类型都可以作为类型参数
func Print[T any](x T) {
	fmt.Print(x)
}

// comparable 只有可比较的类型才可以作为类型参数
// func Max[T comparable](x, y T) T {
// 	if x > y {
// 		return x
// 	}
// 	return y
// }

// // numeric 只有数字类型才可以作为类型参数
// func Sum[T numeric](x, y T) T {
// 	return x + y
// }

// // ordered 只有可排序的类型才可以作为类型参数
// func Min[T ordered](x, y T) T {
// 	if x < y {
// 		return x
// 	}
// 	return y
// }

// // slice 只有切片类型才可以作为类型参数
// func Append[T slice](s []T, x T) []T {
// 	return append(s, x)
// }
