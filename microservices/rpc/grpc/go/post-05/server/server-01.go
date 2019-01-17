package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/liuliqiang/blog-demos/microservices/rpc/grpc/go/post-05/interceptor"
	"github.com/liuliqiang/blog-demos/microservices/rpc/grpc/go/proto-gens"

	"google.golang.org/grpc"
)

var (
	port *int
)

func main() {
	port = new(int)
	flag.IntVar(port, "port", 8080, "port to serve")
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(interceptor.UnaryServerDumpInterceptor()))
	helloworld.RegisterGreeterServer(grpcServer, &helloLiqiangIO{})
	grpcServer.Serve(lis)
}

// hello to https://liqiang.io
type helloLiqiangIO struct {
}

func (*helloLiqiangIO) SayHello(ctx context.Context, req *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{
		Message: fmt.Sprintf("Hello: %s, Welcome to https://liqiang.io", req.Name),
	}, nil
}
