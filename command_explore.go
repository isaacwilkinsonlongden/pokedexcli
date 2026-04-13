package main

import "fmt"

func commandExplore(cfg *Config, areaName string) error {
	pokemonList, err := cfg.pokeapiClient.GetPokemonList(areaName)
	if err != nil {
		return err
	}

	fmt.Println("Exploring " + areaName + "...")
	fmt.Println("Found Pokemon:")
	for _, pokemon := range pokemonList.Results {
		fmt.Println("- " + pokemon.Pokemon.Name)
	}

	return nil
}
