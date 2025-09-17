package main

import (
	"context"
	"fmt"
	"grpc-demo/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	options := grpc.WithTransportCredentials(insecure.NewCredentials())
	clientConn, err := grpc.NewClient("localhost:50051", options)
	if err != nil {
		log.Fatalln(err)
	}

	serviceClient := proto.NewDemoServiceClient(clientConn)

	ctx := context.Background()
	addReq := &proto.AddRequest{
		X: 100,
		Y: 200,
	}
	addRes, err := serviceClient.Add(ctx, addReq)
	if err != nil {
		fmt.Println("error :", err)
		return
	}
	fmt.Printf("Add Result : %d\n", addRes.GetResult())
}
