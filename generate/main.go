package main

import (
	"fmt"
	"minwheel/generate"
)

func main() {
	fmt.Println("go generate")
	var d generate.Weekday
	fmt.Println(d)
	fmt.Println(generate.Weekday((1)))
}
