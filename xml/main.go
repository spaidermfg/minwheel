package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	selectXml()
}

func selectXml() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []string
	for {
		t, err := dec.Token()
		if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}

		if err == io.EOF {
			break
		}

		switch t := t.(type) {
		case xml.StartElement:
			stack = append(stack, t.Name.Local)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				fmt.Printf("%s : %s\n", strings.Join(stack, " "), t)
			}
		}
	}
}

// 输出xml文档中指定元素下的文本
func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}

		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}
