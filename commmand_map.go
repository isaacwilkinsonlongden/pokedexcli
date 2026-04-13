package main

import (
	"fmt"

	"github.com/isaacwilkinsonlongden/pokedexcli/internal/pokeapi"
)

func commandMap(cfg *Config) error {
	locationList, err := pokeapi.GetLocationList(cfg.Next)
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
