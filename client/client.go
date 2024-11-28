package main

import (
	"context"
	"log"
	"time"

	pb "github.com/HADLakmal/go-RPC/protobuff"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("grpc-server:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewCustomCalculatorClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.Calculate(ctx, &pb.Input{Num_1: 4, Num_2: 3})
	if err != nil {
		log.Fatalf("could not calculate: %v", err)
	}
	log.Printf("calculate result: %d", res.Result)

	res, err = client.Multiply(ctx, &pb.Input{Num_1: 4, Num_2: 3})
	if err != nil {
		log.Fatalf("could not multiply: %v", err)
	}
	log.Printf("multiply result: %d", res.Result)
}
