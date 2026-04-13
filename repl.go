package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *Config) error
}

type Config struct {
	Next     *string
	Previous *string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays next map locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous map locations",
			callback:    commandMapb,
		},
	}
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := &Config{}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		userInput := scanner.Text()
		words := cleanInput(userInput)
		if len(words) == 0 {
			continue
		}

		command, ok := getCommands()[words[0]]
		if ok {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(text))
	return words
}
