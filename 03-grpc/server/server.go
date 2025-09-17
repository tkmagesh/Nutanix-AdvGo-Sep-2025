package main

import (
	"context"
	"fmt"
	"grpc-demo/proto"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DemoServiceServerImpl struct {
	proto.UnimplementedDemoServiceServer
}

func (dsi *DemoServiceServerImpl) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	md := ctx.Value("svc-metadata")
	fmt.Println(md)

	x := req.GetX()
	y := req.GetY()
	log.Printf("[service] processing Add x = %d & y = %d\n", x, y)
	time.Sleep(5 * time.Second)
	if ctx.Err() != nil {
		fmt.Println(ctx.Err())
		return nil, nil
	}
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

func (dsi *DemoServiceServerImpl) Aggregate(serverStream proto.DemoService_AggregateServer) error {
	var sum, min, max int64 = 0, 9223372036854775807, -9223372036854775808
LOOP:
	for {
		req, err := serverStream.Recv()
		if err == io.EOF {
			log.Println("[AppService - Aggregate] All the data have been received")
			res := &proto.AggregateResponse{
				Sum: sum,
				Min: min,
				Max: max,
			}
			if err := serverStream.SendAndClose(res); err != io.EOF && err != nil {
				log.Fatalln(err)
			}
			break LOOP
		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(req)
		time.Sleep(2 * time.Second)
		no := req.GetNo()
		sum += no
		if no < min {
			min = no
		}
		if no > max {
			max = no
		}
	}
	return nil
}

func (dsi *DemoServiceServerImpl) Greet(serverStream proto.DemoService_GreetServer) error {
	for {
		greetReq, err := serverStream.Recv()
		if code := status.Code(err); code == codes.Unavailable {
			fmt.Println("Client connection closed")
			break
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		person := greetReq.GetPerson()
		firstName := person.GetFirstName()
		lastName := person.GetLastName()
		log.Printf("Received greet request for %q and %q\n", firstName, lastName)
		message := fmt.Sprintf("Hi %s %s, Have a nice day!", firstName, lastName)
		time.Sleep(2 * time.Second)
		log.Printf("Sending response : %q\n", message)
		greetResp := &proto.GreetResponse{
			Message: message,
		}
		if err := serverStream.Send(greetResp); err != nil {
			if code := status.Code(err); code == codes.Unavailable {
				fmt.Println("Client connection closed")
				break
			}
		}
	}
	return nil
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
