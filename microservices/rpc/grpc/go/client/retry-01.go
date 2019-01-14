package main

import (
	"context"
	"github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"log"

	"github.com/liuliqiang/blog-demos/microservices/rpc/grpc/go/proto-gens"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080",
		grpc.WithStreamInterceptor(grpc_retry.StreamClientInterceptor()),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor()),)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	cli := helloworld.NewGreeterClient(conn)
	resp, err := cli.SayHello(context.Background(), &helloworld.HelloRequest{Name: "lucifer"})
	if err != nil {
		resp, err = cli.SayHello(context.Background(), &helloworld.HelloRequest{Name: "lucifer"})
		if err != nil {
			panic(err)
		}
	}
	log.Printf("[D] resp: %s", resp.Message)
}
