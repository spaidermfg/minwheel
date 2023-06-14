package main

import (
	"log"
	"net/rpc"
)

func main() {
	// client, err := rpc.Dial("tcp", "localhost:7642")
	// if err != nil {
	// 	log.Fatal("rpc dial error:", err)
	// }
	//
	// var reply string
	// //HelloService.Hello：grpc服务名和方法名
	// if err := client.Call(HelloServiceName+".Hello", "mark", &reply); err != nil {
	// 	log.Fatal("rpc connection error:", err)
	// }
	//
	// log.Println(reply)

   client, err := DialHelloService("tcp", "localhost:7642")
   if err != nil {
     log.Fatal("dialing failed", err)
   }
  
   var reply string
   err = client.Hello("mark", &reply)
   if err != nil {
     log.Fatal("hello")
   }
   
   log.Println(reply)
   
  
}

type HelloServiceClient struct {
  *rpc.Client
}

//var _ HelloServiceInterface = (*HelloServiceClient) (nil)

func DialHelloService(network, address string) (*HelloServiceClient, error) {
  c, err := rpc.Dial(network, address)
  if err != nil {
    return nil, err
  }

  return &HelloServiceClient{Client: c}, err
}

func (p *HelloServiceClient) Hello(request string, reply *string) error {
  return p.Client.Call(HelloServiceName+".Hello", request, reply)
}

const HelloServiceName = "path/to/pkg.HelloService"
