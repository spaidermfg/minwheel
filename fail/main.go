package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
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

	isError()

	testMyError()

	err = CopyFile("./fail/a.txt", "./fail/b.txt")
	check(err)
}

func newError() error {
	return errors.New("this is a error")
}

func newWrapError(err error) error {
	return fmt.Errorf("wrap err: %w", err)
}

func isError() {
	err := bufio.ErrBufferFull
	err1 := fmt.Errorf("wrap err1: %w", err)
	err2 := fmt.Errorf("wrap err2: %w", err1)
	if errors.Is(err2, bufio.ErrBufferFull) {
		log.Println(bufio.ErrBufferFull)
	}
}

// 错误处理策略
// 1.透明错误处理策略
// 	 直接return
// 2.哨兵错误处理策略
// 	 使用Go官方定义导出的哨兵错误值
//   eg: bufio.ErrBufferFull
//   使用Is方法将error类型变量与哨兵错误值进行比较
//   if err == bufio.ErrBufferFull 等同于 errors.Is(err, bufio.ErrBufferFull)
//   Is方法会随着错误链向上找到匹配错误值
// 3.错误值类型处理策略
//   使用errors.As()方法,相当于通过类型断言判断一个error类型变量是否为特定的自定义错误类型

type MyError struct {
	e string
}

func (e *MyError) Error() string {
	return e.e
}

func testMyError() {
	err := &MyError{e: "this is wrong"}
	err1 := fmt.Errorf("wrap err1: %w", err)
	err2 := fmt.Errorf("wrap err2: %w", err1)
	var e *MyError
	if errors.As(err2, &e) {
		fmt.Println("MyError is on the chain of err2")
		fmt.Println(errors.Is(e, err), errors.Is(e, err1), errors.Is(e, err2))
		return
	}

	fmt.Println("MyError is not on the chain of err2")
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func CopyFile(src, dst string) (err error) {
	var r, w *os.File

	defer func() {
		if r != nil {
			r.Close()
		}

		if w != nil {
			w.Close()
		}

		if e := recover(); e != nil {
			if w != nil {
				os.Remove(dst)
			}
			err = fmt.Errorf("copy %s %s: %v", src, dst, err)
		}
	}()

	abs, err := filepath.Abs(".")
	check(err)
	log.Println(abs)

	r, err = os.Open(src)
	check(err)

	w, err = os.Create(dst)
	check(err)

	_, err = io.Copy(w, r)
	check(err)

	return nil
}
