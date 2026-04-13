package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *Config, pokemonName string) error {
	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	catchChance := rand.Intn(700)

	if catchChance >= pokemon.BaseExperience {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		cfg.pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}
