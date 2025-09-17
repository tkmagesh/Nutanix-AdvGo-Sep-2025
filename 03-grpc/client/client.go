package main

import (
	"context"
	"fmt"
	"grpc-demo/proto"
	"io"
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

	// doRequestResponse(ctx, serviceClient)
	doServerStreaming(ctx, serviceClient)
}

func doRequestResponse(ctx context.Context, serviceClient proto.DemoServiceClient) {
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

func doServerStreaming(ctx context.Context, serviceClient proto.DemoServiceClient) {
	req := &proto.PrimeRequest{
		Start: 2,
		End:   100,
	}
	clientStream, err := serviceClient.GeneratePrimes(ctx, req)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		res, err := clientStream.Recv()
		if err == io.EOF {
			fmt.Printf("[client] Thats all folks!")
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("[client] Prime No : %d\n", res.GetPrimeNo())
	}

}
