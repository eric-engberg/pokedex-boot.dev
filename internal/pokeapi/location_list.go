package pokeapi

import (
	"encoding/json"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (LocationAreaResponse, error) {
	url := baseURL + "/location-area"
	if pageURL != nil && *pageURL != "" {
		url = *pageURL
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

	locationsResponse := LocationAreaResponse{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&locationsResponse)

	if err != nil {
		return LocationAreaResponse{}, err
	}

	return locationsResponse, nil

}
