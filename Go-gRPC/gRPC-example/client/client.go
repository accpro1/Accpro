package main

import (
	"context"
	"log"

	"mymodule/example" // Generated code

	"google.golang.org/grpc"
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := example.NewGreeterClient(conn)

    response, err := c.SayHello(context.Background(), &example.HelloRequest{Name: "World"})
    if err != nil {
        log.Fatalf("could not greet: %v", err)
    }
    log.Printf("Greeting: %s", response.Message)
}
