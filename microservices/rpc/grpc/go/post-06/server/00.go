package main

import (
	"fmt"
	"log"
	"net"

	"github.com/liuliqiang/blog-demos/microservices/rpc/grpc/go/post-06/proto-gens"
	"google.golang.org/grpc"
)

type server struct {
}

func (*server) Sum(req protobuf.Post06_SumServer) (err error) {
	var reqObj *protobuf.Request

	var sum int64 = 0
	for {
		reqObj, err = req.Recv()
		if err != nil {
			req.SendAndClose(&protobuf.Response{Result: sum})
		} else {
			sum += reqObj.Num
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	protobuf.RegisterPost06Server(grpcServer, &server{})
	grpcServer.Serve(lis)
}
