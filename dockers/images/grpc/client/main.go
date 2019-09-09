package main

import (
	"context"

	helloworld "github.com/liuliqiang/grpc-demo/proto"

	"github.com/liuliqiang/log4go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var cert = "/go/src/github.com/liuliqiang/grpc-demo/90-tls/httpbin.example.com/2_intermediate/certs/ca-chain.cert.pem"

func main() {
	// Create the client TLS credentials
	creds, err := credentials.NewClientTLSFromFile(cert, "")
	if err != nil {
		panic(err)
	}

	// Create a connection with the TLS credentials
	conn, err := grpc.Dial("httpbin.example.com:443", grpc.WithTransportCredentials(creds))
	// conn, err := grpc.Dial("35.239.42.83:80", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer func(conn *grpc.ClientConn) {
		if err = conn.Close(); err != nil {
			log4go.Error(context.Background(), "Failed to close connection: %v", err)
		}
	}(conn)

	cli := helloworld.NewGreeterClient(conn)
	resp, err := cli.SayHello(context.Background(), &helloworld.HelloRequest{Name: "lucifer"})
	if err != nil {
		log4go.Error(context.Background(), "Failed to say hello: %v", err)
	}
	log4go.Info(context.Background(), "resp: %s", resp.Message)
}
