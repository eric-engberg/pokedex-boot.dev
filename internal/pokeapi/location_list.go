package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (LocationAreaResponse, error) {
	url := baseURL + "/location-area?offset=0&limit=20"
	if pageURL != nil && *pageURL != "" {
		url = *pageURL
	}

	data, exists := c.cache.Get(url)
	if exists {
		var locationsResponse LocationAreaResponse
		if err := json.Unmarshal(data, &locationsResponse); err != nil {
			return LocationAreaResponse{}, err
		}
		return locationsResponse, nil

	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return LocationAreaResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	c.cache.Add(url, body)
	locationsResponse := LocationAreaResponse{}
	if err := json.Unmarshal(body, &locationsResponse); err != nil {
		return LocationAreaResponse{}, err
	}

	return locationsResponse, nil
}
