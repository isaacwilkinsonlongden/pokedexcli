package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/isaacwilkinsonlongden/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *Config, areaName string) error
}

type Config struct {
	pokeapiClient pokeapi.Client
	Next          *string
	Previous      *string
	pokedex       map[string]pokeapi.Pokemon
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
		"explore": {
			name:        "explore",
			description: "Displays pokemon found at area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch a Pokemon",
			callback:    commandCatch,
		},
	}
}

func startRepl(cfg *Config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		userInput := scanner.Text()
		words := cleanInput(userInput)
		if len(words) == 0 {
			continue
		}

		var arg string
		if len(words) > 1 {
			arg = words[1]
		}

		command, ok := getCommands()[words[0]]
		if ok {
			err := command.callback(cfg, arg)
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
