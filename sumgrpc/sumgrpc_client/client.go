package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"sumgrpc/sumpb"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	fmt.Println("hello client")

	cc, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer cc.Close()

	c := sumpb.NewSumServiceClient(cc)

	doUnary(c)
}

func doUnary(c sumpb.SumServiceClient) {
	req := &sumpb.SumRequest{
		Data: &sumpb.SumData{
			FirstNumber:  40,
			SecondNumber: 50,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.Sum(ctx, req)
	if err != nil {
		log.Fatalf("error while calling Sum RPC: %v", err)
	}

	log.Printf("Response from Sum: %v", res.Result)
}
