package main

import "fmt"

func main() {
	m := make(map[string]string)
	m["mhz"] = "1000"
	m["sn"] = "ahfasidfhasdf"
	m["version"] = "23232323"
	m["os"] = "dfdf22d"
	m["resettime"] = ""

	for i, v := range m {
		if v == "" {
			fmt.Println(i)
		}
	}

	exists("a")
	exists("b")
	exists("c")
}

var m = map[string]string{"a": "a", "b": "b"}

func exists(str string) {
	if s, ok := m[str]; ok {
		fmt.Println("-------exists", s)
	} else {
		fmt.Println("-------no", ok)
	}
}
