package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemonList(areaName string) (PokemonListResponse, error) {
	requestURL := "https://pokeapi.co/api/v2/location-area/" + areaName

	val, ok := c.cache.Get(requestURL)
	if ok {
		pokemonListResponse := PokemonListResponse{}
		err := json.Unmarshal(val, &pokemonListResponse)
		if err != nil {
			return PokemonListResponse{}, err
		}

		return pokemonListResponse, nil
	}

	res, err := http.Get(requestURL)
	if err != nil {
		return PokemonListResponse{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonListResponse{}, err
	}

	if res.StatusCode > 299 {
		return PokemonListResponse{}, fmt.Errorf("response failed with status code: %d", res.StatusCode)
	}
	c.cache.Add(requestURL, body)

	pokemonListResponse := PokemonListResponse{}
	err = json.Unmarshal(body, &pokemonListResponse)
	if err != nil {
		return PokemonListResponse{}, err
	}

	return pokemonListResponse, nil
}
