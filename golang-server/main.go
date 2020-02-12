package main

import (
	"fmt"
	"github.com/tangxusc/grpc-demo/pkg/test"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

//go:generate protoc test.proto --go_out=plugins=grpc:.
func main() {
	lis, err := net.Listen("tcp", ":9999")
	if err != nil {
		panic(err.Error())
	}
	s := grpc.NewServer()
	test.RegisterHelloServer(s, &HelloService{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type HelloService struct {
}

func (h *HelloService) SayHello(ctx context.Context, request *test.HelloRequest) (*test.HelloReply, error) {
	fmt.Printf("收到客户端请求,name:%s \n", request.Name)
	reply := &test.HelloReply{
		Message: "hello " + request.Name,
	}
	return reply, nil
}
