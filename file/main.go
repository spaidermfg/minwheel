package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// a := "conf/auth.conf"
	// pre, suf := path.Split(a)
	// arr := strings.Split(suf, ".")
	// normalPath := path.Join(pre, arr[0]+"bak"+path.Ext(suf))
	// fmt.Println(normalPath)
	path := "dasserver.conf"
	paths := "dasserverbak.conf"
	file, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	destFile, err := os.OpenFile(paths, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer destFile.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		fmt.Println(line[:1])
		if !strings.Contains(line, "=") && !strings.Contains(line, "#") || 0 == len(line) || line == "\r\n" {
			fmt.Println("Bad line: ", line)
			continue
		} else {
			fmt.Println("Normal line: ", line)
			destFile.WriteString(line)
		}
	}
	os.Rename(paths, path)
}
