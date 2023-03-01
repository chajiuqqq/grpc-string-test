package main

import (
	"grpc-string-test/pb"
	"grpc-string-test/server/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:1234")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("server at 127.0.0.1:1234")
	grpcServer := grpc.NewServer()
	stringService := new(service.StringService)
	pb.RegisterStringServiceServer(grpcServer, stringService)
	grpcServer.Serve(lis)
}
