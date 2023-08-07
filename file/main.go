package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// a := "conf/auth.conf"
	// pre, suf := path.Split(a)
	// arr := strings.Split(suf, ".")
	// normalPath := path.Join(pre, arr[0]+"bak"+path.Ext(suf))
	// fmt.Println(normalPath)
	//findBadLine()
	//changeContent()
	deleteBadLine()
}

func findBadLine() {
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
			if _, err = destFile.WriteString(line); err != nil {
				log.Fatal(err)
			}

		}
	}
	os.Rename(paths, path)
}

func changeContent() {
	filename := "./config.yaml"

	file, err := os.OpenFile(filename, os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err, "open file failed!")
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, isPrefix, err := reader.ReadLine()
		if err == io.EOF || err != nil {
			log.Fatal(err, "end line!")
		}

		log.Println("line:", string(line), "is prefix:", isPrefix)
	}
}

const (
	FILE    = "hello.conf"
	NEWFILE = "hello_new.conf"
)

func deleteBadLine() {
	//以可读写方式打开文件
	file, err := os.OpenFile(FILE, os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//创建一个文件并以可读写方式打开
	destFile, err := os.OpenFile(NEWFILE, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer destFile.Close()

	//将文件内容写入缓冲区中
	reader := bufio.NewReader(file)

	for {
		//以换行符来遍历文件
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		//读到文件结尾
		if err == io.EOF {
			break
		}

		//根据特定条件筛选行
		if len(line) == 0 || line == "\r\n" || !strings.Contains(line, "=") && !strings.Contains(line, "#") || strings.Contains(line, "#") && !strings.HasPrefix(line, "#") {
			log.Println("Bad line:", line)
			continue
		} else {
			log.Println("Normal line:", line)
			//将可用的行写入新文件
			if _, err = destFile.WriteString(line); err != nil {
				log.Fatal(err)
			}
		}
	}
	//处理完毕，将新创建的文件重命名
	//其他系统需要先删除源文件
	os.Rename(NEWFILE, FILE)
}
