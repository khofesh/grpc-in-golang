package main

import (
	"calculator/calculatorpb"
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	calculatorpb.UnimplementedSumServiceServer
}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Sum function was invoked with %v", req)

	firstNumber := req.GetData().GetFirstNumber()
	secondNumber := req.GetData().GetSecondNumber()
	result := firstNumber + secondNumber

	res := &calculatorpb.SumResponse{
		Result: result,
	}

	return res, nil
}

func main() {
	fmt.Println("hello from server")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	calculatorpb.RegisterSumServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
