package main

import (
	"bufio"
	"encoding/binary"
	"encoding/gob"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// 使用gob读取抽象数据类型
// 读取抽象数据类型
// 1. fmt.Fscan & fmt.Fprint
// 2. binary.Read & binary.Write
// 3. gob.Decode & gob.Encode

func main() {
	// a := "conf/auth.conf"
	// pre, suf := path.Split(a)
	// arr := strings.Split(suf, ".")
	// normalPath := path.Join(pre, arr[0]+"bak"+path.Ext(suf))
	// fmt.Println(normalPath)
	//findBadLine()
	//changeContent()
	//deleteBadLine()

	//useDirect()

	useAbstractWrite()

	catchWrite()

	catchRead()
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

func directWriteBytesToFile(path string, data []byte) (int, error) {
	if path == "" {
		return 0, errors.New("path is null")
	}

	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND|os.O_SYNC, 0666)
	if err != nil {
		return 0, err
	}

	defer func() {
		file.Close()
		file.Sync()
	}()

	return file.Write(data)
}

func directReadBytesFromFile(path string, data []byte) (int, error) {
	if path == "" {
		return 0, errors.New("path is null")
	}

	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	return file.Read(data)
}

func useDirect() {
	file := "./foo.txt"
	text := "Find basic SPL concepts in this post: SPL concepts for beginners. For beginners, \n" +
		"you can find characteristic basic computations of SPL in SPL Operations for Beginners. \n" +
		"Experienced programmers can quickly understand the differences between SPL and SQL. \n" +
		"A software architect can understand the differences between SPL and traditional databases after reading Q&A of esProc Architecture. \n"
	n, err := directWriteBytesToFile(file, []byte(text))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("write %d bytes to %s\n", n, file)

	buf := make([]byte, 38)
	_, err = directReadBytesFromFile(file, buf)
	if err != nil {
		log.Fatal(err)
	}

	open, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer open.Close()

	for {
		content, err1 := open.Read(buf)
		if err1 == io.EOF {
			fmt.Println("read meets EOF")
			break
		}

		if err1 != nil {
			fmt.Println("read file error:", err1)
			return
		}

		fmt.Printf("read %d bytes from file: %q\n", content, buf)

	}
}

// 写入抽象数据类型

type Player struct {
	name   string
	age    int
	gender string
}

func (p Player) String() string {
	return fmt.Sprintf("%v %v %v", p.name, p.gender, p.age)
}

func directWriteToFile(path string, players []*Player) error {
	create, err := os.Create(path)
	if err != nil {
		return err
	}

	defer func() {
		create.Sync()
		create.Close()
	}()

	for _, player := range players {
		// 写
		_, err = fmt.Fprintf(create, "%s\n", player)
		if err != nil {
			return err
		}

	}

	return err
}

func useAbstractWrite() {
	players := []*Player{
		{"tom", 22, "man"},
		{"mark", 23, "women"},
		{"dylan", 24, "man"},
		{"frank", 25, "women"},
	}

	err := directWriteToFile("./ab.txt", players)
	if err != nil {
		log.Fatal(err)
	}

	player, err := backToAb("./ab.txt")
	if err != nil {
		log.Fatal(err)
	}

	for _, ab := range player {
		fmt.Println(ab.name, ab.gender, ab.age)
	}

	//binary
	useBinary()

	//gob
	useGob()
}

func backToAb(path string) ([]*Player, error) {
	open, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer open.Close()

	players := make([]*Player, 0)
	//player := &Player{}

	for {
		var player Player
		// 读
		_, err = fmt.Fscanf(open, "%s %s %d", &player.name, &player.gender, &player.age)
		if err == io.EOF {
			fmt.Println("read meet EOF")
			return players, nil
		}

		if err != nil {
			fmt.Println("read file error:", err)
			return nil, err
		}

		players = append(players, &player)
	}

	//return players, err
}

///  ---------------------------binary read & write

type Student struct {
	Name   [20]byte
	Age    int8
	Gender [6]byte
}

func binaryWriteToFile(path string, students []Student) error {
	create, err := os.Create(path)
	if err != nil {
		return err
	}
	defer create.Close()

	for _, v := range students {
		if err = binary.Write(create, binary.BigEndian, &v); err != nil {
			return err
		}
	}

	return nil
}

func binaryReadFromFile(path string) ([]*Student, error) {
	open, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer open.Close()

	students := make([]*Student, 0)

	for {
		var student Student
		// 大端字节序
		err = binary.Read(open, binary.BigEndian, &student)
		if err == io.EOF {
			fmt.Println("read meet EOF")
			return students, nil
		}

		if err != nil {
			return nil, err
		}

		students = append(students, &student)
	}
}

func useBinary() {
	var students [3]Student

	copy(students[0].Name[:], "Tommy")
	copy(students[1].Name[:], "Mark")
	copy(students[2].Name[:], "Rank")

	copy(students[0].Gender[:], "man")
	copy(students[1].Gender[:], "women")
	copy(students[2].Gender[:], "man")

	students[0].Age = 18
	students[1].Age = 19
	students[2].Age = 10

	err := binaryWriteToFile("./bin.txt", students[:])
	if err != nil {
		log.Fatal(err)
	}

	stus, err := binaryReadFromFile("./bin.txt")
	if err != nil {
		log.Fatal(err)
	}

	for _, stu := range stus {
		fmt.Printf("--- %s %s %d\n", stu.Name, stu.Gender, stu.Age)
	}
}

// gob Decode & Encode

type Portrait struct {
	Light     string
	Color     int8
	Structure string
}

func gobWriteToFile(path string, portrait []*Portrait) error {
	create, err := os.Create(path)
	if err != nil {
		return err
	}
	defer create.Close()

	encoder := gob.NewEncoder(create)
	for _, p := range portrait {
		if err = encoder.Encode(p); err != nil {
			return err
		}
	}

	return err
}

func useGob() {
	portraits := []*Portrait{
		{"open", 89, "二八"},
		{"close", 34, "三八"},
		{"close", 56, "curve"},
	}

	err := gobWriteToFile("./gob.txt", portraits)
	if err != nil {
		log.Fatal(err)
	}

	open, err := os.Open("./gob.txt")
	if err != nil {
		log.Fatal(err)
	}

	var portrait Portrait
	dec := gob.NewDecoder(open)
	for {
		err = dec.Decode(&portrait)
		if err == io.EOF {
			fmt.Println("read meet EOF")
			break
		}

		if err != nil {
			log.Fatal("read file error: ", err)
		}

		fmt.Println(portrait)
	}
}

// 使用带缓冲I/O模式
func catchWrite() {
	file := "./bufio.txt"
	f, err := os.OpenFile(file, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		f.Sync()
		f.Close()
	}()

	data := []byte("I love golang!\n")

	// 创建带缓冲的类型
	bio := bufio.NewWriterSize(f, 32)

	// 写入15字节，存入缓存区
	bio.Write(data)

	// 写入15字节，存入缓存区
	bio.Write(data)

	// 写入15字节，向文件中写入32字节，存入 （3 * 15）- 32字节
	bio.Write(data)

	bio.Flush()
}

func catchRead() {
	file := "./bufio.txt"
	f, err := os.Open(file)
	if err != nil {
		log.Fatal("open file error:", err)
	}

	// 通过包裹函数创建带缓冲I/O的类型
	// 初始缓冲区大小危64字节
	bio := bufio.NewReaderSize(f, 64)
	fmt.Printf("初始缓冲区缓冲数量：%v\n", bio.Buffered())

	i := 1
	for {
		data := make([]byte, 15)
		n, err1 := bio.Read(data)
		if err1 == io.EOF {
			fmt.Printf("第%v次读取数据，读取到文件末尾，程序退出\n", i)
			break
		}

		if err1 != nil {
			log.Fatal("read file error:", err1)
		}

		fmt.Printf("第%d次读取数据：%q,长度为%v\n\n", i, data, n)
		fmt.Printf("当前缓冲区缓存数据量为%d\n", bio.Buffered())
		i++
	}
}
