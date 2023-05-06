# 《Go语言底层原理刨析读书笔记》

***

## 1. 深入了解Go语言编译器 go compiler

编译器一般分为三个阶段，第一个阶段编译器前端主要专注于理解源程序、扫描解析源程序并进行精准的语义表达。编译器的中间阶段会对代码进行多次优化，例如识别冗余代码、识别内存逃逸等。编译器后端主要主要用于生成特定目标机器上的程序，例如汇编语言、obj等。

**词法解析**

词法解析过程会扫描Go源文件，并将代码进行符号化。相关代码文件位于src/cmd/compile/internal目录下。另外，Go语言标准库go/scanner和go/token中也提供了用于扫描源代码的接口。

模拟编译器词法解析过程

``` golang
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
```

运行结果

``` bash
1:1     IDENT   "cos"
1:4     (       ""
1:5     IDENT   "x"
1:6     )       ""
1:8     +       ""
1:10    IMAG    "2i"
1:12    *       ""
1:13    IDENT   "sin"
1:16    (       ""
1:17    IDENT   "x"
1:18    )       ""
1:20    COMMENT "// Euler"
1:28    ;       "\n"
```

**语法解析**

Go语言采用了标准的自上而下的递归下降算法，一个Go源文件中主要有包导入声明，静态常量，类型声明，变量声明以及函数声明。函数声明是文件中最复杂的一类语法，每一种声明语法或表达式都有对应的结构体。    
核心的算法位于syntax/nodes.go和syntax/parser.go中。	

**抽象语法树构建**

抽象语法树(Abstarct Syntax Tree)是一种常见的树状结构的中间态, 在Go中，任何一种import，type，const，func都是一个根节点，在根节点下包含当前声明的子节点。    
核心逻辑代码位于gc/noder.go和gc/syntax.go中。

**类型检查**

在完成语法树的构建后，便会遍历节点树来决定节点的类型。这其中包括代码中明确指定的类型，例如`var a int`, 和隐式类型，例如`a := 1`。    
还会对一些语法和语义进行检查，比如引用的结构体名称是否为大写可导出，数组字面量访问是否超过长度，计算编译时常量等相关操作。    
核心的处理代码位于gc/typecheck.go中。    

**变量捕获**

