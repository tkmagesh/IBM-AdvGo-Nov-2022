package services

import (
	"context"
	"errors"
	"fmt"
	"grpc-app/proto"
	"log"
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

func (asi *AppServiceImpl) GeneratePrimes(req *proto.PrimeRequest, serverStream proto.AppService_GeneratePrimesServer) error {
	start := req.GetStart()
	end := req.GetEnd()
	fmt.Printf("Generating prime number between %d and %d\n", start, end)
	for no := start; no <= end; no++ {
		if isPrime(no) {
			fmt.Printf("Sending Prime No : %d\n", no)
			res := &proto.PrimeResponse{
				PrimeNo: no,
			}
			err := serverStream.Send(res)
			if err != nil {
				log.Fatalln(err)
			}
			time.Sleep(500 * time.Millisecond)
		}
	}
	fmt.Println("All prime numbers are sent")
	return nil
}

func isPrime(no int32) bool {
	for i := int32(2); i < (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
