package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"

	pb "github.com/saufiroja/microservice-grpc/calculate/proto"
	"google.golang.org/grpc"
)

type Server struct {
	pb.CalculateServiceServer
}

var addr string = "localhost:50052"

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Listening on %s", addr)

	s := grpc.NewServer()
	pb.RegisterCalculateServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *Server) Calculate(ctx context.Context, req *pb.CalculateRequest) (*pb.CalculateResponse, error) {
	log.Printf("Calculate function was invoked with %v", req)

	return &pb.CalculateResponse{
		Result: req.Num1 + req.Num2,
	}, nil
}

func (s *Server) Primes(req *pb.PrimesRequest, stream pb.CalculateService_PrimesServer) error {
	log.Printf("Primes function was invoked with %v", req)

	var k int32 = 2
	var N int32 = req.GetNum()

	for N > 1 {
		if N%k == 0 {
			res := &pb.PrimesResponse{
				Result: k,
			}
			stream.Send(res)
			N = N / k
		} else {
			k++
		}
	}

	return nil
}

func (s *Server) Avg(stream pb.CalculateService_AvgServer) error {
	log.Printf("Avg function was invoked")

	var res int32

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				Result: res,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		log.Printf("Receiving %v\n", req)
	}
}

func (s *Server) Max(stream pb.CalculateService_MaxServer) error {
	log.Printf("Max function was invoked")
	var max int32

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Printf("Error while reading client streaming: %v\n", err)
		}

		number := req.Num
		fmt.Println("num", number)

		// NUMBER > MAX
		// 	1 > 0 true
		// 	5 > 1 true
		// 	3 > 5 false
		// 	6 > 5 true
		// 	2 > 6 false
		// 	10 > 6 true
		if number > max {
			max = number
			err := stream.Send(&pb.AvgResponse{
				Result: max,
			})

			if err != nil {
				log.Fatalf("Error while sending data to client: %v\n", err)
			}
		}
	}
}
