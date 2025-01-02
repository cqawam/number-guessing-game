package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)
var reader = bufio.NewReader(os.Stdin)
func welcome() {
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I am thinking of a number between 1 and 100.")
	fmt.Println("You have 5 chances to guess the correct number.")
	fmt.Print("\n")

	fmt.Println("Please select the difficulty level:")
	fmt.Println("1. Easy (10 chances)")
	fmt.Println("2. Medium (5 chances)")
	fmt.Println("3. Hard (3 chances)")
	fmt.Print("\n")
}

func game(chances int) {
	randomNumber := rand.Intn(100) + 1

	// Loop for the number of chances according to level
	for i := 1; i <= chances; i++ {
		fmt.Print("Enter your guess: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal()
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal()
		}

		if guess > randomNumber {
			fmt.Printf("Incorrect! The number is less than %d\n\n", guess)
			continue
		} else if guess < randomNumber {
			fmt.Printf("Incorrect! The number is greater than %d\n\n", guess)
			continue
		} else {
			fmt.Printf("Congratulations! You guessed the correct number in %d attempts\n", i)
			break
		}
	}
	fmt.Printf("Sorry, you're out of chances! The correct number was %d.\n", randomNumber)
}

func main() {
	welcome()

	// Get the difficulty level from console input
	fmt.Print("Enter your choice: ")
	level, _, err := reader.ReadRune()
	if err != nil {
		log.Fatal()
	}
	// Discard leftover input
	reader.ReadString('\n')

	switch level {
	case '1':
		fmt.Println("Great! You have selected the Easy difficulty level.")
		fmt.Print("Let's start the game!\n\n")
		game(10)
	case '2':
		fmt.Println("Great! You have selected the Medium difficulty level.")
		fmt.Print("Let's start the game!\n\n")
		game(5)
	case '3':
		fmt.Println("Great! You have selected the Hard difficulty level.")
		fmt.Print("Let's start the game!\n\n")
		game(3)
	default:
		fmt.Println("Incorrect difficulty level")
	}

	// Implement the game function with each difficulty.
}
