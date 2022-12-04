package main

import (
	"context"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

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
	// doAvg(c)
	// doMax(c)
	doSqrt(c, -10)

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

func doMax(c pb.CalculateServiceClient) {
	log.Println("doMax was invoked")

	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	req := []*pb.MaxRequest{
		{Num: 1},
		{Num: 5},
		{Num: 3},
		{Num: 6},
		{Num: 2},
		{Num: 10},
	}

	ch := make(chan struct{})

	go func() {
		for i, v := range req {
			log.Printf("Send number: %v\n", req[i])
			stream.Send(v)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while receiving: %v\n", err)
				break
			}

			log.Printf("Received a new maximum: %v\n", res.Result)
		}

		close(ch)
	}()

	<-ch
}

func doSqrt(c pb.CalculateServiceClient, n int32) {
	log.Println("doSqrt was invoked")

	req := &pb.SqrtRequest{
		Num: n,
	}

	res, err := c.Sqrt(context.Background(), req)

	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			log.Printf("Error message from server: %s\n", e.Message())
			log.Printf("Error code from server: %s\n", e.Code())

			if e.Code() == codes.InvalidArgument {
				log.Println("We probably sent a negative number!")
				return
			}
		} else {
			log.Fatalf("A non gRPC error: %v\n", err)
		}
	}

	log.Printf("Sqrt: %f\n", res.Result)

}
