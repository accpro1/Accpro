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

    player1_win := []string{}
    player2_win := []string{}
    
	fmt.Print("------Players started playing the game-------\n ")
    
    moves := []string{"rock", "paper", "scissors", "rock", "paper"}
    
    for _, move := range moves {
        fmt.Println("Move played by Player 1:", move)
 
        req := &rps.MoveRequest{
            PlayerId: "Player 1",
             Move: move, 
        }
   
      ctx, cancel := context.WithTimeout(context.Background(), time.Second)
      defer cancel()

       res, err := client.PlayGame(ctx, req)
       if err != nil {
          log.Fatalf("Failed to play game: %v", err)
        }

      log.Printf("\n Result of the rounds : %s  %s", res.GetWinner(), res.GetMessage())

      if res.GetWinner() == "Player 1" {
         player1_win = append(player1_win, "win")
        } else if res.GetWinner() == "Player 2"{
            player2_win = append(player2_win, "win")
        } else {
            fmt.Println("Nothing to add on any side, its tie move")
        }
    }

    //Final Winner calculation
    if len(player1_win) > len(player2_win)    {
     fmt.Println("Final Game winner is Player 1 !!!")
    } else if len(player1_win) == len(player2_win) {
        fmt.Println("Final Game winner is Tie between Player 1 and Player 2 !!!")
    } else {
        fmt.Println("Final Game winner is Player 2 !!!")
    }
}