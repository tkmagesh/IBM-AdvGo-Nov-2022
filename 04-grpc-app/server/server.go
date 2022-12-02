package main

import (
	"grpc-app/proto"
	"log"
	"net"

	"grpc-app/server/services"

	"google.golang.org/grpc"
)

func main() {
	asi := &services.AppServiceImpl{}
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterAppServiceServer(grpcServer, asi)
	grpcServer.Serve(listener)
}
