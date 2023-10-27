package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

// encoding/json

var origin = []byte(`{
	"name": "mark",
	"age": 40,
	"address": "La"
}`)

// 压缩json，使json格式化变得更紧凑
func TestCompact(t *testing.T) {
	b := bytes.Clone(origin)
	fmt.Printf("Formatted json:\n	%v\n", string(b))

	var out bytes.Buffer
	if err := json.Compact(&out, b); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Compact json:\n		%v\n", string(out.Bytes()))

	//	output:
	//	Formatted json:
	//		{
	//			"name": "mark",
	//			"age": 40,
	//			"address": "La"
	//		}
	//	Compact json:
	//		{"name":"mark","age":40,"address":"La"}

}

func TestJson(t *testing.T) {

}

//
//func TestJson(t *testing.T) {
//
//}
//
//func TestJson(t *testing.T) {
//
//}
//
//func TestJson(t *testing.T) {
//
//}
//
//func TestJson(t *testing.T) {
//
//}
//
//func TestJson(t *testing.T) {
//
//}
//
//func TestJson(t *testing.T) {
//
//}
//
//func TestJson(t *testing.T) {
//
//}
