package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocationExplore(exploreLocation string) (LocationExplore, error) {
	url := baseURL + "/location-area" + "/" + exploreLocation

	if val, ok := c.cache.Get(url); ok {
		exploreResp := LocationExplore{}
		err := json.Unmarshal(val, &exploreResp)
		if err != nil {
			return LocationExplore{}, err
		}

		return exploreResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationExplore{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationExplore{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationExplore{}, err
	}

	exploreResp := LocationExplore{}
	err = json.Unmarshal(data, &exploreResp)
	if err != nil {
		return LocationExplore{}, err
	}

	c.cache.Add(url, data)
	return exploreResp, nil
}
