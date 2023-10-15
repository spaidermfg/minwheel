package main

import "fmt"

func main() {
	printBit()
}

func printBit() {
	i := 67
	fmt.Printf("十进制：%d,二进制：%b\n", i, i)

	i = 67 << 2
	fmt.Printf("十进制：%d,二进制：%b\n", i, i)

}
