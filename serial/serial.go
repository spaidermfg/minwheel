package main

import (
	serial1 "go.bug.st/serial"
	"log"

	_ "github.com/tarm/serial"

	_ "github.com/jacobsa/go-serial/serial"
)

// 串口通信
// 常用repo
//  https://github.com/bugst/go-serial.git   => go get go.bug.st/serial
// 	https://github.com/jacobsa/go-serial.git => go get github.com/jacobsa/go-serial/serial
// 	https://github.com/tarm/serial.git       => go get github.com/tarm/serial

func main() {
	ports, err := serial1.GetPortsList()
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
