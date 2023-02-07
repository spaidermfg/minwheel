package main

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"
	"unicode/utf16"
)

// go 1.20新版本功能测试
// https://go.dev/doc/go1.20

func main() {
	StringsCollect()
	UnicodeCollect()
	TimeCollect()
	FilePathCollect()

}

// The new CutPrefix and CutSuffix functions are like TrimPrefix and TrimSuffix but also report whether the string was trimmed.
// The new Clone function allocates a copy of a string.
func StringsCollect() {
	fmt.Println("=========strings=========")
	name := "hello world"
	after, founds := strings.CutPrefix(name, "h")
	fmt.Println(after, founds) //ello world true
	before, found := strings.CutSuffix(name, "rld")
	fmt.Println(before, found)       //hello wo true
	fmt.Println(strings.Clone(name)) // hello world;  s == name  (true)
	fmt.Println("=========================")
}

// The new AppendRune function appends the UTF-16 encoding of a given rune to a uint16 slice,
// analogous to utf8.AppendRune.
func UnicodeCollect() {
	fmt.Println("=========unicode=========")
	var a []uint16
	a = append(a, 3)
	a = append(a, 4)
	u := utf16.AppendRune(a, 2)
	fmt.Println(u) //[3 4 2]
	fmt.Println("=========================")
}

// The new time layout constants DateTime, DateOnly,
// and TimeOnly provide names for three of the most common layout strings used in a survey of public Go source code.

// The new Time.Compare method compares two times.

// Parse now ignores sub-nanosecond precision in its input, instead of reporting those digits as an error.

// The Time.MarshalJSON method is now more strict about adherence to RFC 3339.
func TimeCollect() {
	fmt.Println("========time=============")
	fmt.Println(time.DateTime) // 2006-01-02 15:04:05
	fmt.Println(time.DateOnly) // 2006-01-02

	t := time.Unix(time.Now().Unix()+7*60*60*24, 0) // 2023-02-09 11:10:19 +0800 CST
	u := time.Now()                                 // 2023-02-02 11:10:19.959169 +0800 CST
	fmt.Println(t, u, t.Compare(u))                 // 1

	b, _ := time.Now().MarshalJSON()
	fmt.Println(b)
	fmt.Println("=========================")
}

// The new error SkipAll terminates a Walk immediately but successfully.
// The new IsLocal function reports whether a path is lexically local to a directory.

// For example, if IsLocal(p) is true,
// then Open(p) will refer to a file that is lexically within the subtree rooted at the current directory.
func FilePathCollect() {
	fmt.Println("=========filepath========")
	path1 := "/Users/mac/Downloads/self.doc"
	a := filepath.IsLocal(path1)
	fmt.Println(a) //false
	path2 := "mac/Downloads/self.doc"
	b := filepath.IsLocal(path2)
	fmt.Println(b) //true
	fmt.Println("=========================")
}

func MimeCollect() {

}
