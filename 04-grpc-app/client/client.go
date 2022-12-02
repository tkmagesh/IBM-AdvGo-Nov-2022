package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	options := grpc.WithTransportCredentials(insecure.NewCredentials())
	clientConn, err := grpc.Dial("localhost:50051", options)
	if err != nil {
		log.Fatalln(err)
	}
	client := proto.NewAppServiceClient(clientConn)
	rootCtx := context.Background()

	//doRequestResponse(rootCtx, client)
	doServerStreaming(rootCtx, client)
}

func doRequestResponse(ctx context.Context, client proto.AppServiceClient) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	addRequest := &proto.AddRequest{
		X: 100,
		Y: 200,
	}
	addResponse, err := client.Add(timeoutCtx, addRequest)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Result : %d\n", addResponse.GetResult())
}

func doServerStreaming(ctx context.Context, client proto.AppServiceClient) {
	req := &proto.PrimeRequest{
		Start: 3,
		End:   100,
	}
	clientStream, err := client.GeneratePrimes(ctx, req)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		res, err := clientStream.Recv()
		if err == io.EOF {
			fmt.Println("All prime numbers received")
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("Prime No : %d\n", res.GetPrimeNo())
	}
}
