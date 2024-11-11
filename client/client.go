package main

import (
	"context"

	proto "github.com/go-grpc-unary-operation/protoc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:9091", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(err)
	}

	client := proto.NewExampleClient(conn)

	req := &proto.HelloRequest{SomeString: "Hellow from the boss Shrishail: Hi"}

	client.ServerReply(context.TODO(), req)
}
