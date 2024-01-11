package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	fmt.Println("This a game where you guess the number, I am thinking.")
	//Generate a random number between 0 and 100
	var randomNumber int = rand.Intn(100)
	// fmt.Println(randomNumber)
	//Number of tries
	var numberOfTries = 0
	var input int
	//Grab the input of the user
	fmt.Scan(&input)
	numberOfTries++

	//While the number does not equal to the random number
	for randomNumber != input {
		if 0 > input || input > 100 {
			fmt.Println("The number needs to be Between 0 and 100:")
			fmt.Scan(&input)
		}
		if 0 <= input && input < randomNumber {
			fmt.Println("You are lower from the number")

			fmt.Scan(&input)
			numberOfTries++
		}
		if input > randomNumber && input <= 100 {
			fmt.Println("You are high from the number")

			fmt.Scan(&input)
			numberOfTries++
		}
	}
	fmt.Printf("The number of tries %d \n", numberOfTries)

	var continueOn string
	//To continue playing again
	fmt.Println("Would you like to play again? Y or N")
	fmt.Scan(&continueOn)
	if strings.ToLower(continueOn) == "y" {
		main()
	}
}
