package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		cleanedInput := cleanInput(userInput)
		firstWord := cleanedInput[0]

		fmt.Printf("Your command was: %s\n", firstWord)
	}
}
