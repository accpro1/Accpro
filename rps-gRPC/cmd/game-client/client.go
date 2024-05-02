// client1.go
package main

import (
	"context"
	"log"
	"mymodule/rps"
	"time"

	"fmt"

	"google.golang.org/grpc"
)

func main() {
    // Establish a gRPC connection to the server
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("Failed to connect: %v", err)
    }
    defer conn.Close()

    // Create an RPS client from the connection
    client := rps.NewRPSClient(conn)
    
	fmt.Print("----Player 1 is going to play for move 1---\n ")
	fmt.Print("Choose any of the option from rock, paper, or scissors: ")
	var user_choice string
	fmt.Scanln(&user_choice )
	fmt.Print("\n You choose: ", user_choice)

    // Example request to play a game session
    req := &rps.MoveRequest{
        PlayerId: "Player 1",
        Move: user_choice, // Can be "rock", "paper", or "scissors"
    }

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    res, err := client.PlayGame(ctx, req)
    if err != nil {
        log.Fatalf("Failed to play game: %v", err)
    }

    log.Printf("\n Game result: %s  %s", res.GetWinner(), res.GetMessage())
}
