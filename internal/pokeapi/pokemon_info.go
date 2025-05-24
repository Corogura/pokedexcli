package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemonInfo(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name

	if val, exists := c.cache.Get(url); exists {
		infoResp := Pokemon{}
		err := json.Unmarshal(val, &infoResp)
		if err != nil {
			return Pokemon{}, err
		}
		return infoResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	infoResp := Pokemon{}
	err = json.Unmarshal(dat, &infoResp)
	if err != nil {
		return Pokemon{}, err
	}
	return infoResp, nil
}
