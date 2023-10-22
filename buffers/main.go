package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// 提供高效的缓冲I/O操作
func main() {
	input := "Hello bufio.\nThis is a new line\n"
	reader := bufio.NewReader(strings.NewReader(input))

	for {
		//readString, err := reader.ReadString('\n')
		//if err == io.EOF {
		//	fmt.Println("EOF")
		//	break
		//}
		//if err != nil {
		//	log.Fatal(err)
		//}
		//fmt.Printf(readString)

		//readByte, err := reader.ReadByte()
		//if err != nil {
		//	log.Fatal(err)
		//}
		//fmt.Println(readByte)
		//fmt.Println(string(readByte))

		//bytes, err := reader.ReadBytes('\n')
		//if err != nil {
		//	log.Fatal(err)
		//}
		//fmt.Println(bytes)
		//fmt.Println(string(bytes))

		//line, prefix, err := reader.ReadLine()
		//if err != nil {
		//	log.Fatal(err)
		//}
		//fmt.Printf("line: %v, string(line): %v, prefix: %v\n", line, string(line), prefix)

		line, err := reader.ReadSlice('\n')
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v, %v", line, string(line))
	}

	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString("Hello, ")
	writer.WriteString("bufio.\n")
	writer.Flush()

	dir, _ := os.Getwd()
	fmt.Println(dir)
	open, err := os.OpenFile("a.txt", os.O_RDWR, 0777)
	if err != nil {
		log.Fatal(err)
	}
	defer open.Close()
	newReader := bufio.NewReader(open)
	for {
		readString, err1 := newReader.ReadString('\t')
		if err1 == io.EOF {
			break
		}
		fmt.Print(readString)
	}

	//open.WriteString("This is a old world\n\tGive some music.")
	newWriter := bufio.NewWriter(open)
	writeString, err := newWriter.WriteString("This is a new world\nThank your for my future.\n")
	err2 := writer.Flush()
	fmt.Println(writeString, err, err2)
}
