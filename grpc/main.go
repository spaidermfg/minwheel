package main

import (
	"io"
	"log"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// RPC 和 Protobuf
// Remote Procedure Call 远程过程调用,  net/rpc
// RPC规则：方法只能有两个可序列化的参数，其中第二个参数是指针类型，并返回一个error类型，必须是公开的方法。

func main() {
	//将对象类型中所有满足rpc规则的方法注册为rpc函数
	//	rpc.RegisterName("HelloService", new(HelloService))
	log.Println("grpc server start...")
	RegisterHelloService(new(HelloService))

	//base of http protocol
	http.HandleFunc("httprpc", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: r.Body,
			Writer:     w,
		}
		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})

	http.ListenAndServe(":7642", nil)
	//listener, err := net.Listen("tcp", ":7642")
	//if err != nil {
	//	log.Fatal("Listen tcp error:", err)
	//}
	//
	//for {
	//	conn, err := listener.Accept()
	//	if err != nil {
	//		log.Fatal("accept error:", err)
	//	}
	//
	//	//在tcp链接上提供rpc服务
	//	//go rpc.ServeConn(conn)
	//	go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	//}
}

type HelloService struct{}

func (h *HelloService) Hello(request string, reply *string) error {
	*reply = "hello: " + request
	return nil
}

const HelloServiceName = "path/to/pkg.HelloService"

type HelloServiceInterface interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}
