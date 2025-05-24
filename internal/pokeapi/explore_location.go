package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ExploreLocation(location string) (RespPokemons, error) {
	url := baseURL + "/location-area/" + location

	if val, exists := c.cache.Get(url); exists {
		pokemonResp := RespPokemons{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return RespPokemons{}, err
		}
		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemons{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemons{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokemons{}, err
	}

	pokemonResp := RespPokemons{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return RespPokemons{}, err
	}
	return pokemonResp, nil
}
