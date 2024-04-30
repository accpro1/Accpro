package main

import (
	"context"
	"log"
	"net"

	"mymodule/example" // Generated code

	"google.golang.org/grpc"
)

type server struct{
    example.UnimplementedGreeterServer  // Embedding the struct
}

func (s *server) SayHello(ctx context.Context, req *example.HelloRequest) (*example.HelloResponse, error) {
    return &example.HelloResponse{Message: "Hello, " + req.Name + "!"}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    example.RegisterGreeterServer(s, &server{})
    log.Println("Server listening on port 50051")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
