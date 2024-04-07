package main

import (
	"fmt"
	"golang.org/x/net/html/charset"
	"log"
	"os"
	"unicode/utf8"
)

func main() {
	detectFormat()
}

func detectFormat() {
	file, err := os.ReadFile("noodlesss.csv")
	if err != nil {
		log.Fatal(err)
	}

	files, err := os.ReadFile("noodless.csv")
	if err != nil {
		log.Fatal(err)
	}
	e, name, certain := charset.DetermineEncoding(file, "")
	fmt.Println(e, name, certain)

	valid := utf8.Valid(file)
	valids := utf8.Valid(files)
	fmt.Println(valid, valids)
}
