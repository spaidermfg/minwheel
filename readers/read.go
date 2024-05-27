package main

import (
	"io"
	"log"
	"strings"
)

// 统计字符数量
func countLetters(r io.Reader) (map[string]int, error) {
	buf := make([]byte, 2048)
	out := make(map[string]int)

	for {
		n, err := r.Read(buf)
		for _, v := range buf[:n] {
			if (v > 'a' && v < 'z') || (v > 'A' && v < 'Z') {
				out[string(v)]++
			}
		}

		if err == io.EOF {
			return out, nil
		}

		if err != nil {
			return nil, err
		}
	}
}

func main() {
	str := "Hello, Is there a coffee on the table."
	reader := strings.NewReader(str)
	count, err := countLetters(reader)
	log.Println(count, err)
}
