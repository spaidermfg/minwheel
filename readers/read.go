package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"
)

// 统计字符数量
func countLetters(r io.Reader) (map[string]int, error) {
	buf := make([]byte, 100)
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

type person struct {
	Name string
	Age  int8
}

func jsonEn(p person) {
	temp, err := os.CreateTemp(".", "*.json")
	if err != nil {
		log.Fatal(err)
	}
	defer temp.Close()

	err = json.NewEncoder(temp).Encode(p)
	if err != nil {
		log.Fatal(err)
	}
}

func jsonDe() {

}

func main() {
	str := "Hello, Is there a coffee on the table."
	reader := strings.NewReader(str)
	count, err := countLetters(reader)
	log.Println(count, err)

	// json.Encode
	p := person{
		Name: "mark",
		Age:  23,
	}
	jsonEn(p)
}
