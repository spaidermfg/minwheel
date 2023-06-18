package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"grpc/protobuf"
	"io"
	"log"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// RPC 和 Protobuf
// Remote Procedure Call 远程过程调用,  net/rpc
// RPC规则：方法只能有两个可序列化的参数，其中第二个参数是指针类型，并返回一个error类型，必须是公开的方法。
// protobuf
// 最基本的数据单元： message
// 通过成员的唯一编号来绑定对应的数据
// 生成proto文件相对应的go代码需要使用protoc工具，还需要安装针对go语言的代码生成插件protoc-gen-go
// grpc是基于protobuf开发的跨语言的开源rpc框架， 基于http2.0协议设计

func main() {
	//将对象类型中所有满足rpc规则的方法注册为rpc函数
	//	rpc.RegisterName("HelloService", new(HelloService))
	log.Println("grpc server start...")
	//tls加密
	creds, err := credentials.NewServerTLSFromFile("grpc/server.crt", "grpc/server.key")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer(grpc.Creds(creds))

	//启动反射服务
	reflection.Register(s)
	//protobuf.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))
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

func (h *HelloService) HelloProtoBuf(request *protobuf.String, reply *protobuf.String) error {
	reply.Value = "Hello: " + request.GetValue()
	return nil
}

type HelloServiceImpl struct {
}

func (h *HelloServiceImpl) mustEmbedUnimplementedHelloServiceServer() {
	//TODO implement me
	panic("implement me")
}

// 基于服务器端的grpc
func (h *HelloServiceImpl) HelloProtobuf(ctx context.Context, args *protobuf.String) (*protobuf.String, error) {
	reply := &protobuf.String{Value: "Hello: " + args.GetValue()}
	return reply, nil
}

// 接收客户端发来的消息
func (h *HelloServiceImpl) Channel(stream protobuf.HelloService_ChannelServer) error {
	for {
		args, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		reply := &protobuf.String{
			Value: "Hello:" + args.GetValue(),
		}
		err = stream.Send(reply)
		if err != nil {
			return err
		}
	}
}
