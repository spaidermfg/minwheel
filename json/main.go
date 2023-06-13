package main

import (
	"encoding/json"
	"fmt"
	"log"

)

func main() {
  str := `{"code":-1,"count":1,"data":"P5 添加失败 UNIQUE constraint failed: point.gn","msg":"error"}`

  out := make(map[string]interface{})
  if err := json.Unmarshal([]byte(str), &out); err != nil {
    log.Fatal(err)
  }

  fmt.Println(out["data"])
  fmt.Println(out["count"])
  fmt.Println(out["data"])
  fmt.Println(out["msg"])
}
