package main

import (
	"context"
	"google.golang.org/grpc"
	"net"
	"log"

)

type HelloServiceImpl struct {
}

func (h *HelloServiceImpl) Hello(ctx context.Context, args *String) (*String, error) {
	reply := &String{Value: "Hello:" + args.GetValue()}
	return reply, nil
}

func main() {
	grpcServer := grpc.NewServer()
	RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))
	listen, err := net.Listen("tcp", ":2333")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(listen)
}
