package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
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
	timeoutCtx, cancel := context.WithTimeout(rootCtx, 3*time.Second)
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
