package client

import (
	"context"
	"fmt"
	"github.com/tangxusc/grpc-demo/pkg/test"
	"google.golang.org/grpc"
	"testing"
)

func TestClient(t *testing.T) {
	dial, err := grpc.Dial("localhost:9999", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	defer dial.Close()
	client := test.NewHelloClient(dial)
	reply, err := client.SayHello(context.TODO(), &test.HelloRequest{
		Name: "test1",
	})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("收到服务器回复:%s \n", reply.Message)
}
