package main

import (
	"context"
	"grpc-app/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type appServiceImpl struct {
	proto.UnimplementedAppServiceServer
}

func (asi *appServiceImpl) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	x := req.GetX()
	y := req.GetY()

	result := x + y

	res := &proto.AddResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	asi := &appServiceImpl{}
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterAppServiceServer(grpcServer, asi)
	grpcServer.Serve(listener)
}
