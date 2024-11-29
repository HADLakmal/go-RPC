# gRPC Microservice Communication

This project demonstrates the use of **gRPC** for efficient and high-performance communication between microservices. gRPC is a modern Remote Procedure Call (RPC) framework that supports load balancing, tracing, health checking, authentication, and distributed computing.

## Features of gRPC

- **High Performance**: Efficient serialization using Protocol Buffers.
- **Abstracted Communication**: Simplifies client-server interaction by behaving like local function calls.
- **Error Handling**: Returns errors or exceptions depending on the programming language.
- **Scalability**: Suitable for internal communication in backend services.
- **API Contracts**: Easily defined using Interface Description Language (IDL) such as Protocol Buffers.
- **Streaming Support**: Supports single or bidirectional streaming over a single connection.

## Protobuf Definition

The API contract is defined in a `.proto` file as follows:

```proto
syntax = "proto3";

package myservice;

option go_package = ".";

service CustomCalculator {
    rpc Calculate (Input) returns (Output);
    rpc Multiply (Input) returns (Output);
}

message Input {
    int32 num_1 = 1;
    int32 num_2 = 2;
}

message Output {
    int32 result = 1;
}
```

This service provides two functions:

1. **Calculate**: Takes two integers and returns their sum.
2. **Multiply**: Takes two integers and returns their product.

## Generating gRPC Code

Install the `protoc` compiler and run the following command to generate gRPC stubs:

```bash
protoc --go_out=. --go-grpc_out=. contract.proto
```

This will generate two `.pb.go` files containing the server and client stubs.

## Server Implementation

The server is responsible for executing the gRPC functions. Below is an example implementation in Go:

```go
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
```

Run the server with:

```bash
go mod tidy
go run server.go
```

## Client Implementation

The client connects to the server and invokes the gRPC functions. Below is an example implementation:

```go
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
 conn, err := grpc.Dial("grpc-server:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
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
```

Run the client with:

```bash
go mod tidy
go run client.go
```

## Output

The client will display the following output:

```plaintext
calculate result: 7
multiply result: 12
```

## Docker Support

The codebase includes Docker implementations to enable seamless deployment in distributed environments. Refer to the GitHub repository for more details.

## Repository

Find the full code and additional details in the [GitHub Repository](https://github.com/HADLakmal/go-RPC).

