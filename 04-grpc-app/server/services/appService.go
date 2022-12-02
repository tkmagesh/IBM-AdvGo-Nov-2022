package services

import (
	"context"
	"fmt"
	"grpc-app/proto"
)

type AppServiceImpl struct {
	proto.UnimplementedAppServiceServer
}

func (asi *AppServiceImpl) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	x := req.GetX()
	y := req.GetY()
	fmt.Printf("Add Request received, x = %d and y = %d\n", x, y)
	result := x + y

	res := &proto.AddResponse{
		Result: result,
	}
	return res, nil
}
