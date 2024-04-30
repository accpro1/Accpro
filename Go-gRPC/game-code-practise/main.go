package main

import (
	"fmt"
	"math/rand"
)


func getRandomString(arr []string) string {
    index := rand.Intn(len(arr))
    return arr[index]
}

func main2() {
	//user choice
	fmt.Print("Choose any of the option from rock, paper, or scissors: ")
    var user_choice string
	fmt.Scanln(&user_choice )
	fmt.Print("\n You choose: ", user_choice)

    // computer choice
	var options = []string{"rock", "paper", "scissors"}
	computer_choice := getRandomString(options)
	fmt.Print("\n Computer choose: ", computer_choice)

    //logic of the game
	if user_choice == computer_choice {
        fmt.Println("\n It's a tie! ")
    } else if user_choice == "rock" && computer_choice == "scissors" {
        fmt.Println("\n You win! ")
    } else if user_choice == "paper" && computer_choice == "rock" {
        fmt.Println("\n You win! ")
    } else if user_choice == "scissors" && computer_choice == "paper" {
        fmt.Println("\n You win! ")
    } else {
		fmt.Println("\n Computer win! ")
	}
    

}