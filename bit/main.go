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

	i = i | 0000000011
	fmt.Printf("十进制：%d,二进制：%b\n", i, i)

	k := 10
	k = k << 3
	k = k + (10 << 1)

	fmt.Printf("十进制：%d,二进制：%b\n", k, k)
}
