package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationList(url *string) (LocationListResponse, error) {
	requestURL := "https://pokeapi.co/api/v2/location-area/"
	if url != nil {
		requestURL = *url
	}

	res, err := http.Get(requestURL)
	if err != nil {
		return LocationListResponse{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationListResponse{}, err
	}

	if res.StatusCode > 299 {
		return LocationListResponse{}, fmt.Errorf("response failed with status code: %d", res.StatusCode)
	}

	locationListResponse := LocationListResponse{}
	err = json.Unmarshal(body, &locationListResponse)
	if err != nil {
		return LocationListResponse{}, err
	}

	return locationListResponse, nil
}
