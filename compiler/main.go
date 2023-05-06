package main

import (
	"fmt"
	"go/scanner"
	"go/token"
)

// go 编译器 词法解析
func main() {
	//表达式，模拟为一个文件
	src := []byte("cos(x) + 2i*sin(x) // Euler")

	var scan scanner.Scanner
	fs := token.NewFileSet()
	file := fs.AddFile("", fs.Base(), len(src))
	scan.Init(file, src, nil, scanner.ScanComments)

	for {
		pos, tok, lit := scan.Scan()
		if tok == token.EOF {
			break
		}

		fmt.Printf("%s\t%s\t%q\n", fs.Position(pos), tok, lit)
	}
}
