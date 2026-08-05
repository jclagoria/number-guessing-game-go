package main

import (
	"bufio"
	"fmt"
	"number-guessing-game-cli/internal/game"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	welcome()
	d, ok := chooseDifficulty(reader)
	if !ok {
		return
	}

	round := game.New(game.RandomSecret(), d)
	fmt.Println("Let's start the game!")

	play(reader, round)
}

func welcome() {
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println("Please select the difficulty level:")
	fmt.Println("1. Easy (10 chances)")
	fmt.Println("2. Medium (5 chances)")
	fmt.Println("3. Hard (3 chances)")
}

func chooseDifficulty(reader *bufio.Reader) (game.Difficulty, bool) {
	for {
		n, ok, eof := readInt(reader, "Enter you choice: ", 1, 3)
		if eof {
			return 0, false
		}

		if !ok {
			fmt.Println("Invalid choice. Please enter 1, 2, or 3.")
			continue
		}

		d := game.Difficulty(n - 1)
		fmt.Println("Great! You have selected the %s difficulty level.\n", d.String())
		return d, false
	}
}

func play(reader *bufio.Reader, round *game.Round) {
	for !round.Over() {
		n, ok, eof := readInt(reader, "Enter your guess", 1, 100)
		if eof {
			return
		}
		if !ok {
			fmt.Println("Invalid guess. Plase enter a number between 1 and 100")
			continue
		}

		switch round.Guess(n) {
		case game.Correct:
			fmt.Printf("Congratulations! You guesses the correct number in %d attemps.\n", round.Attempts())
		case game.TooHigh:
			fmt.Printf("Incorrect! The number is less than %d.\n", n)
		case game.TooLow:
			fmt.Printf("Incorrect! Tjhe number is greater than %d", n)
		}

		if !round.Won() {
			fmt.Printf("You lost! The number was %d.\n", round.Secret())
		}
	}
}

// readInt reads a line, trims it, and parses it as an integer in [min, max].
// Returns (n, true, false) for a valid int, (0, false, false) for invalid input,
// and (0, false, true) on EOF.
func readInt(reader *bufio.Reader, prompt string, min, max int) (int, bool, bool) {
	fmt.Print(prompt)

	line, err := reader.ReadString('\n')
	if err != nil {
		return 0, false, true
	}

	n, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil || n < min || n > max {
		return 0, false, false
	}

	return n, true, false
}
