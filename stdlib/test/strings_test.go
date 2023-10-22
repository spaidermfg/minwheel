package test

import (
	"fmt"
	"strings"
	"testing"
	"unsafe"
)

// strings

func TestClone(t *testing.T) {
	str := "Beautiful"
	clone := strings.Clone(str)
	fmt.Printf("str: %v, %p\nclone: %v, %p\n", str, &str, clone, &clone)
	fmt.Printf("sizeof: %v,%v\n", unsafe.Sizeof(str), unsafe.Sizeof(clone))
	fmt.Printf("byte: %v,%v\n", []byte(str), []byte(clone))
}

// 比较的是首字母的ASCII码
// a == b, return 0
// a < b, return -1
// a > b, return 1
func TestCompare(t *testing.T) {
	a := "Beautiful"
	b := "sky"
	compare := strings.Compare(a, b)
	fmt.Println("compare:", compare)
	fmt.Printf("byte: %v,%v\n", []byte(a), []byte(b))
}

// strings.Index的调用
// 子字符串必须连续
func TestContains(t *testing.T) {
	a := "Beautiful"
	b := "ti"
	contains := strings.Contains(a, b)
	fmt.Println("a contains b?", contains)
	fmt.Printf("byte: %v,%v\n", []byte(a), []byte(b))
}

// 只要有一个字符匹配，就返回true
func TestContainsAny(t *testing.T) {
	a := "Beautiful"
	b := "hjkl"
	containsAny := strings.ContainsAny(a, b)
	fmt.Println("a containsAny b?", containsAny)
	fmt.Printf("byte: %v,%v\n", []byte(a), []byte(b))
}

func TestContainsFunc(t *testing.T) {
	a := "Beautiful"
	containsFunc := strings.ContainsFunc(a, func(r rune) bool {
		if a == "Beautiful" {
			return true
		}
		return false
	})
	fmt.Println("a containsFunc b?", containsFunc)
}

//func TestC(t *testing.T) {
//	strings.Compare()
//}

//func TestC(t *testing.T) {
//	strings.Compare()
//}

//func TestC(t *testing.T) {
//	strings.Compare()
//}
