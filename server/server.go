package main

import (
	"context"
	"log"
	"net"

	pb "github.com/HADLakmal/go-RPC/protobuff"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCustomCalculatorServer
}

func (s *server) Calculate(ctx context.Context, req *pb.Input) (*pb.Output, error) {
	return &pb.Output{Result: req.Num_1 + req.Num_2}, nil
}

func (s *server) Multiply(ctx context.Context, req *pb.Input) (*pb.Output, error) {
	return &pb.Output{Result: req.Num_1 * req.Num_2}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCustomCalculatorServer(grpcServer, &server{})
	log.Println("Server is running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
