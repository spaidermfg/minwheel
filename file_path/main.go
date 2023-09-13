package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	path := "/a/b/c/red (1).csv"
	fmt.Println(filepath.Ext(path))

	fmt.Println(filepath.Base(path))
}
