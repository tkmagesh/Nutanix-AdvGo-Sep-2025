package main

import (
	"context"
	"grpc-demo/proto"
	"log"
	"net"

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
