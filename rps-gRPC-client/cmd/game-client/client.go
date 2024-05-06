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

    counter_1 := 0
    counter_2 := 0
    
	fmt.Print("\n------Players started playing the game-------\n ")
    
    moves := []string{"rock", "paper", "scissors", "rock", "paper"}
    
    for index, move := range moves {
        fmt.Println("\n Move played by Player 1:", move)
 
        req := &rps.MoveRequest{
            PlayerId: "Player 1",
             Move: move, 
        }
   
      ctx, cancel := context.WithTimeout(context.Background(), time.Second)
      defer cancel()

       res, err := client.PlayGame(ctx, req)
       if err != nil {
          log.Fatalf("\n Failed to play game: %v", err)
        }

      log.Printf("\n Result of the round %d : %s  %s", index+1, res.GetWinner(), res.GetMessage())

      if res.GetWinner() == "Player 1" {
        counter_1++
        } else if res.GetWinner() == "Player 2"{
           counter_2++
        } else {
            fmt.Println("\n Nothing to add on any side, its tie move")
        }
    }

    //Final Winner calculation
    if counter_1 > counter_2    {
        fmt.Printf("\n Final Game winner is Player 1 !!! \n")
        fmt.Printf("\n Game Over \n")
       } else if counter_1 == counter_2 {
           fmt.Printf("\n Final Game winner is Tie between Player 1 and Player 2 !!! \n")
           fmt.Printf("\n Game Over \n")
       } else {
           fmt.Printf("\n Final Game winner is Player 2 !!! \n")
           fmt.Printf("\n Game Over \n")
       }
}