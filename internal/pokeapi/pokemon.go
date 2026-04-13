package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	requestURL := "https://pokeapi.co/api/v2/pokemon/" + pokemonName

	val, ok := c.cache.Get(requestURL)
	if ok {
		pokemon := Pokemon{}
		err := json.Unmarshal(val, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}

		return pokemon, nil
	}

	res, err := http.Get(requestURL)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	if res.StatusCode > 299 {
		return Pokemon{}, fmt.Errorf("response failed with status code: %d", res.StatusCode)
	}
	c.cache.Add(requestURL, body)

	pokemon := Pokemon{}
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}
