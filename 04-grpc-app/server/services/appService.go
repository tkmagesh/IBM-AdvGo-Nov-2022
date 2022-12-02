package services

import (
	"context"
	"errors"
	"fmt"
	"grpc-app/proto"
	"io"
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

func (asi *AppServiceImpl) CalculateAverage(serverStream proto.AppService_CalculateAverageServer) error {
	sum := int32(0)
	count := int32(0)
	for {
		req, err := serverStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		no := req.GetNo()
		fmt.Printf("Received No : %d\n", no)
		sum += no
		count++
	}
	fmt.Println("Received all the nos. Sending the response....")
	avg := sum / count
	res := &proto.AverageResponse{
		Average: avg,
	}
	err := serverStream.SendAndClose(res)
	if err != nil {
		log.Fatalln(err)
	}
	return nil
}

//bidirectional streaming
func (asi *AppServiceImpl) Greet(serverStream proto.AppService_GreetServer) error {
	for {
		req, err := serverStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		person := req.GetPerson()
		msg := fmt.Sprintf("Hi %s %s!", person.GetFirstName(), person.GetLastName())
		resp := &proto.GreetResponse{
			GreetMessage: msg,
		}
		time.Sleep(500 * time.Millisecond)
		e := serverStream.Send(resp)
		if e != nil {
			log.Fatalln(err)
		}
	}
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
