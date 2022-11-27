package main

import (
	"context"
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
