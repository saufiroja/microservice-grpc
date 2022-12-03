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

	// doCalculate(c)
	// doPrimes(c)
	doAvg(c)

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

func doAvg(c pb.CalculateServiceClient) {
	log.Println("doAvg was invoked")

	req := []*pb.AvgRequest{
		{Num: 1},
		{Num: 2},
		{Num: 3},
		{Num: 4},
	}

	stream, err := c.Avg(context.Background())
	if err != nil {
		log.Fatalf("Error while calling Avg %v\n", err)
	}

	var res float32
	var angka int32
	for _, v := range req {
		res += float32(v.Num)
		angka = v.Num

	}

	res /= float32(angka)

	result, err := stream.CloseAndRecv()
	result.Result = int32(res)
	if err != nil {
		log.Fatalf("Error while receiving response from avg: %v\n", err)
	}

	log.Println("Avg: ", res)
}
