package main

import (
	"fmt"

	"golang.org/x/net/html/charset"
)

//判断文件编码
func main() {

	// var a interface{}
	// a = 1.669172061407931e+09

	// fmt.Printf("%T\n", a)
	// switch v := a.(type) {
	// case int32:
	// 	fmt.Print(v)
	// case float64:
	// 	fmt.Println(float64(v))
	// 	fmt.Println(int64(v))
	// default:
	// 	fmt.Println(v, "=")
	// }

	// fmt.Println("===========1")
	// b := 1.669172061407931e+09
	// d := fmt.Sprintf("%1.0f", b*1000)
	// fmt.Println(b * 10)
	// fmt.Printf("%1.0f\n", b*1e3) //时间戳转换，float64类型毫秒级时间戳乘以1e3转换为int64类型时间戳
	// fmt.Printf("%T", d)
	// fmt.Println("===========2")
	// var c int64
	// c = 1669173516 * 1e3
	// fmt.Printf("%T, %d\n", c, c)
	// fmt.Println(time.UnixMilli(c * 1e3).String())

	utf8 := []byte{230, 136, 145, 230, 152, 175, 85, 84, 70, 56}
	//gbk := []byte{206, 210, 202, 199, 71, 66, 75}
	encoding, name, certain := charset.DetermineEncoding(utf8, "text/html")
	fmt.Printf("编码：%v\n名称：%s\n确定：%t\n", encoding, name, certain)

}
