package main

import (
	"log"
	"os"
)

func main() {
	args := os.Args
	src := args[1]
	dest := args[2]

	log.Println(os.Getwd())
	log.Println(src, "----------->", dest)
	err := os.Rename(src, dest)
	if err != nil {
		log.Fatal(err)
	}
}
