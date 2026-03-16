package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocation(locationName string) (LocationAreaDetailResponse, error) {
	url := baseURL + "/location-area/" + locationName

	data, exists := c.cache.Get(url)
	if exists {
		var locationDetailResponse LocationAreaDetailResponse
		if err := json.Unmarshal(data, &locationDetailResponse); err != nil {
			return LocationAreaDetailResponse{}, err
		}
		return locationDetailResponse, nil

	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return LocationAreaDetailResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaDetailResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaDetailResponse{}, err
	}
	c.cache.Add(url, body)
	locationDetailResponse := LocationAreaDetailResponse{}
	if err := json.Unmarshal(body, &locationDetailResponse); err != nil {
		return LocationAreaDetailResponse{}, err
	}

	return locationDetailResponse, nil
}
