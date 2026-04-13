package main

import (
	"fmt"
)

func commandMap(cfg *Config) error {
	locationList, err := cfg.pokeapiClient.GetLocationList(cfg.Next)
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

func commandMapb(cfg *Config) error {
	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	locationList, err := cfg.pokeapiClient.GetLocationList(cfg.Previous)
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
