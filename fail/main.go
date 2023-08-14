package main

import (
	"errors"
	"fmt"
)

// 如何优雅的处理错误
// go认为错误就是值，错误的处理就是基于值的决策
// 不要心存侥幸，要始终处理错误
// 错误是以error接口变量的形式
//
//	type error interface {
//		 Error() string
//	}
//
// 构造错误值的两种方法
// errors.New("fail")
// fmt.Errorf("%v", "fail")
//
// errors 和 fmt 返回的实例不同, 都会返回这个：
//
//	type errorString struct {
//			s string
//	}
//
// 但fmt使用%w格式化字符时会返回：
//
//	type wrapError struct {
//			msg string
//			err error
//	}
//
// wrapError实现了Unwrap方法，被其包装的错误可以被检视到
func main() {
	err := newError()
	if err != nil {
		fmt.Println(err)
	}

	err1 := newWrapError(err)
	is := errors.Is(err1, err)
	fmt.Println(is, err1)
}

func newError() error {
	return errors.New("this is a error")
}

func newWrapError(err error) error {
	return fmt.Errorf("wrap err: %w", err)
}
