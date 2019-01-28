package main

import (
	"context"
	"log"

	"github.com/liuliqiang/blog-demos/microservices/rpc/grpc/go/post-06/proto-gens"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	cli := protobuf.NewPost06Client(conn)
	sumCli, err := cli.Sum(context.Background())
	if err != nil {
		panic(err)
	}
	sumCli.Send(&protobuf.Request{Num: int64(1)})
	sumCli.Send(&protobuf.Request{Num: int64(2)})
	sumCli.Send(&protobuf.Request{Num: int64(3)})
	resp, err := sumCli.CloseAndRecv()
	if err != nil {
		panic(err)
	}
	log.Printf("[D] resp: %s", resp.Result)
}
