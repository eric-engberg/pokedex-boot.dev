package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	pokemon := Pokemon{}

	data, exists := c.cache.Get(url)
	if exists {
		if err := json.Unmarshal(data, &pokemon); err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	if err := json.Unmarshal(body, &pokemon); err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, body)

	return pokemon, nil
}
