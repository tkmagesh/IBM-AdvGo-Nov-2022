package services

import (
	"context"
	"errors"
	"fmt"
	"grpc-app/proto"
	"time"
)

type AppServiceImpl struct {
	proto.UnimplementedAppServiceServer
}

func (asi *AppServiceImpl) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("time elapsed : %ds\n", i+1)
		select {
		case <-ctx.Done():
			fmt.Println("Cancel signal received")
			return &proto.AddResponse{}, errors.New("cancel signal received")
		default:
		}
	}
	fmt.Println("Prcessing the request")
	x := req.GetX()
	y := req.GetY()
	fmt.Printf("Add Request received, x = %d and y = %d\n", x, y)
	result := x + y

	res := &proto.AddResponse{
		Result: result,
	}
	return res, nil
}
