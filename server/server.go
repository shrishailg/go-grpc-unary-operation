package main

import (
	"context"
	"errors"
	"fmt"
	"net"

	proto "github.com/go-grpc-unary-operation/protoc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedExampleServer
}

func main() {
	listener, tcpErr := net.Listen("tcp", ":9091")
	if tcpErr != nil {
		panic(tcpErr)
	}

	srv := grpc.NewServer()
	proto.RegisterExampleServer(srv, &server{})
	reflection.Register(srv)

	if err := srv.Serve(listener); err != nil {
		panic(err)
	}
}

func (s *server) ServerReply(ctx context.Context, request *proto.HelloRequest) (*proto.HelloResponse, error) {
	fmt.Println("Recived request from client ", request.SomeString)
	fmt.Println("Hello from the server")

	return &proto.HelloResponse{}, errors.New(" ")
}
