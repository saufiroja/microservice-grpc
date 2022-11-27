package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/saufiroja/microservice-grpc/calculate/proto"
)

func main() {
	conn, err := grpc.Dial(":50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	defer conn.Close()

	c := pb.NewCalculateServiceClient(conn)

	doCalculate(c)
	doPrimes(c)

	log.Printf("Connected to %s", ":50052")
}

func doCalculate(c pb.CalculateServiceClient) {
	req := &pb.CalculateRequest{
		Num1: 10,
		Num2: 20,
	}

	res, err := c.Calculate(context.Background(), req)
	if err != nil {
		log.Fatalf("failed to calculate: %v", err)
	}

	log.Printf("Result: %d", res.Result)
}

func doPrimes(c pb.CalculateServiceClient) {
	req := &pb.PrimesRequest{
		Num: 120,
	}

	stream, err := c.Primes(context.Background(), req)
	if err != nil {
		log.Fatalf("failed to primes: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err != nil {
			break
		}

		log.Printf("Result: %d", res.Result)
	}
}
