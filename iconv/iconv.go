package main

import (
	"github.com/djimenez/iconv-go"
	"log"
)

func main() {
	println("iconv.....")
	// output, err := iconv.ConvertString("asdfasdfasdfasdf", "ISO-IR-6", "UTF-8")
	output, err := iconv.ConvertString("asdfasdfasdfasdf", "ISO-8859-1", "UTF-8")
	if err != nil {
		log.Fatal("error:", err)
	}

	println("result:", output)
}
