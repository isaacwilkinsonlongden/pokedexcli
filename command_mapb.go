package main

import (
	"fmt"

	"github.com/isaacwilkinsonlongden/pokedexcli/internal/pokeapi"
)

func commandMapb(cfg *Config) error {
	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	locationList, err := pokeapi.GetLocationList(cfg.Previous)
	if err != nil {
		return err
	}

	cfg.Next = locationList.Next
	cfg.Previous = locationList.Previous

	for _, location := range locationList.Results {
		fmt.Println(location.Name)
	}

	return nil
}
