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
}
