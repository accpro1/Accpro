// server.go
package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"mymodule/rps"
	"net"

	"google.golang.org/grpc"
)

type RPSServer struct {
    rps.UnimplementedRPSServer // Embeds a default implementation
}

func getRandomString(arr []string) string {
	index := rand.Intn(len(arr))
	return arr[index]
}

func (s *RPSServer) PlayGame(ctx context.Context, req *rps.MoveRequest) (*rps.GameResult, error) {
  
    move1 := req.GetMove()
    player_id := req.GetPlayerId()

	fmt.Print("----Player 2 is going to play for move 2---\n ")
	    // computer choice Player 2
	var options = []string{"rock", "paper", "scissors"}
	computer_choice := getRandomString(options)
	fmt.Print("\n Player 2 choose: ", computer_choice)

	
    //logic of the game
	if move1 == computer_choice {
		 fmt.Println("\n It's a tie!")
		 return &rps.GameResult{Winner: "It's a tie", Message: "!!!"}, nil 
    } else if move1 == "rock" && computer_choice == "scissors" {
		 fmt.Println("\n Game Won by: ", player_id )
		 return &rps.GameResult{Winner: player_id, Message: "Won the Game"}, nil 
    } else if move1 == "paper" && computer_choice == "rock" {
        fmt.Println("\n Game Won by: ", player_id )
		 return &rps.GameResult{Winner: player_id, Message: "Won the Game"}, nil
    } else if move1 == "scissors" && computer_choice == "paper" {
         fmt.Println("\n Game Won by: ", player_id )
		 return &rps.GameResult{Winner: player_id, Message: "Won the Game"}, nil
    } else {
		 fmt.Println("\n Game Won by: ", player_id )
		return &rps.GameResult{Winner: "Player 2", Message: "Won the Game"}, nil
	}
	
}

func main() {
    // Listen on port 50051
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    s := grpc.NewServer()

    rps.RegisterRPSServer(s, &RPSServer{})

    log.Printf("Server is listening on port 50051")

    if err := s.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}