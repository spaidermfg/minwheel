package main

import (
	"log"
	"net"
	"net/rpc"
)

// RPC 和 Protobuf
// Remote Procedure Call 远程过程调用,  net/rpc
// RPC规则：方法只能有两个可序列化的参数，其中第二个参数是指针类型，并返回一个error类型，必须是公开的方法。

func main() {
	//将对象类型中所有满足rpc规则的方法注册为rpc函数
	rpc.RegisterName("HelloService", new(HelloService))
	listener, err := net.Listen("tcp", ":7642")
	if err != nil {
		log.Fatal("Listen tcp error:", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("accept error:", err)
	}

	//在tcp链接上提供rpc服务
	rpc.ServeConn(conn)
}

type HelloService struct{}

func (h *HelloService) Hello(request string, reply *string) error {
	*reply = "hello: " + request
	return nil
}
