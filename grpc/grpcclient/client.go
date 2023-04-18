package main

import (
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:7642")
	if err != nil {
		log.Fatal("rpc dial error:", err)
	}

	var reply string
	//HelloService.Hello：grpc服务名和方法名
	if err := client.Call("HelloService.Hello", "mark", &reply); err != nil {
		log.Fatal("rpc connection error:", err)
	}

	log.Println(reply)
}
