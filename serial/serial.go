package main

import (
	"log"

	serials "go.bug.st/serial"
)

// 串口通信
// 常用repo
//  https://github.com/bugst/go-serial.git   -> go get go.bug.st/serial
// 	https://github.com/jacobsa/go-serial.git
// 	https://github.com/tarm/serial.git

func main() {
	ports, err := serials.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}

	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
	}

	for k, v := range ports {
		log.Println(k, v)
	}
}
