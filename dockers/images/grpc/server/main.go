package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	helloworld "github.com/liuliqiang/grpc-demo/proto"

	"github.com/liuliqiang/log4go"
	"google.golang.org/grpc"
)

var port int

func main() {
	flag.IntVar(&port, "port", 9321, "port to serve")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	helloworld.RegisterGreeterServer(grpcServer, &helloServer{})
	log4go.Info(context.Background(), "Ready to start server @:%d...", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to server grpc: %v", err)
	}
}

type helloServer struct {
}

func (s *helloServer) SayHello(ctx context.Context, req *helloworld.HelloRequest) (resp *helloworld.HelloReply, err error) {
	log4go.Info(context.Background(), "%s say hello to you!", req.Name)
	return &helloworld.HelloReply{
		Message: req.Name,
	}, nil
}
