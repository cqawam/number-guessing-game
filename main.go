package main

import (
	"fmt"
	"math/rand"
)

func game(chances int, level string) {
	fmt.Printf("Great! You have selected the %s difficulty level.\nLet's start the game!\n\n", level)
	randomNumber := rand.Intn(100) + 1
	success := false
outer:
	// Loop for the number of chances according to level
	for i := 1; i <= chances; i++ {
		var guess int
		fmt.Print("Enter your guess: ")
		fmt.Scanf("%d", &guess)

		switch {
		case guess == randomNumber:
			success = true
			fmt.Printf("Congratulations! You guessed the correct number in %d attempts\n", i)
			break outer
		case guess > randomNumber:
			fmt.Printf("Incorrect! The number is less than %d\n\n", guess)
			continue
		case guess < randomNumber:
			fmt.Printf("Incorrect! The number is greater than %d\n\n", guess)
			continue
		}

	}
	if !success {
		fmt.Printf("Sorry, you're out of chances! The correct number was %d.\n", randomNumber)
	}
}

func main() {
	welcome := `Welcome to the Number Guessing Game!
I am thinking of a number between 1 and 100.
You have a number of chances to guess the correct number.

Please select the difficulty level:
1. Easy (10 chances)
2. Medium (5 chances)
3. Hard (3 chances)

`
	fmt.Println(welcome)
	// Get the difficulty level from console input
	var level int
	fmt.Print("Enter your choice: ")
	fmt.Scanf("%d", &level)
	switch level {
	case 1:
		game(10, "Easy")
	case 2:
		game(5, "Medium")
	case 3:
		game(3, "Hard")
	default:
		fmt.Println("Incorrect difficulty level")
	}
}
