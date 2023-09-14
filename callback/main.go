package main

import "fmt"

type Book struct {
	Author string
	Edge   int
}

type Callback func(book []*Book)

// 回调函数
func main() {
	applyCallback("mark", 678, func(book []*Book) {
		for _, v := range book {
			fmt.Println(v)
		}
	})

}

func applyCallback(value string, edge int, cb Callback) {
	books := make([]*Book, 0)
	for i := 0; i < 100; i++ {
		book := &Book{
			Author: fmt.Sprintf("%v-%v", value, i),
			Edge:   i + edge,
		}
		books = append(books, book)
	}

	cb(books)
}
