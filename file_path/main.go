package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	path := "red (1).csv"
	fmt.Println(filepath.Ext(path))
}
