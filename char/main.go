package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	str := '-'
	strs := '-'
	fmt.Print(str)
	fmt.Print(strs)
	a := &User{
		Name: "ddf",
		Age:  22,
	}
	b := &User{
		Name: "ddf",
		Age:  22,
	}
	c := &User{
		Name: "ddf",
		Age:  22,
	}

	u := make([]*User, 0)
	u = append(u, a)
	u = append(u, b)
	u = append(u, c)

	fmt.Println(len(u))
	for k, v := range u {
		fmt.Println(k, v)

	}
	err := os.Chmod("/Users/mac/Downloads/sql.csv", 0777)
	if err != nil {
		fmt.Println(err)
	}

	gg := "aB1_*$%#"
	s := strings.ToLower(gg)
	fmt.Println(s)

	pn := "W3.P1.HOST"
	s2 := strings.SplitAfterN(pn, ".", 3)
	s3 := strings.Split(pn, ".")
	fmt.Println(s2)
	fmt.Println(s3[2])
}

type User struct {
	Name string
	Age  int
}
