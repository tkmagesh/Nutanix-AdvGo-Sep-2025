package main

import (
	"context"
	"fmt"
	"grpc-demo/proto"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type DemoServiceServerImpl struct {
	proto.UnimplementedDemoServiceServer
}

func (dsi *DemoServiceServerImpl) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {

	x := req.GetX()
	y := req.GetY()
	log.Printf("[service] processing Add x = %d & y = %d\n", x, y)
	result := x + y
	res := &proto.AddResponse{
		Result: result,
	}
	return res, nil
}

func (dsi *DemoServiceServerImpl) GeneratePrimes(req *proto.PrimeRequest, serverStream proto.DemoService_GeneratePrimesServer) error {
	start := req.GetStart()
	end := req.GetEnd()
	log.Printf("[service] GeneratePrimes start = %d, end = %d\n", start, end)
	for no := start; no <= end; no++ {
		if isPrime(no) {
			res := &proto.PrimeResponse{
				PrimeNo: no,
			}
			fmt.Printf("[service] sending prime no - %d\n", no)
			if err := serverStream.Send(res); err != nil {
				log.Fatalln(err)
			}
			time.Sleep(500 * time.Millisecond)
		}
	}
	fmt.Printf("[service] All prime numbers are generated!")
	return nil
}

func isPrime(no int64) bool {
	for i := int64(2); i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	dsi := &DemoServiceServerImpl{}
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterDemoServiceServer(grpcServer, dsi)
	grpcServer.Serve(listener)
}
