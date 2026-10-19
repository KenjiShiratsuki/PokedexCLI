package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL string) (Locations, error) {
	url := baseURL + "/location-area?offset=0&limit=20"
	if pageURL != "" {
		url = pageURL
	}
	if dat, ok := c.cache.Get(pageURL); ok {
		locationsResp := Locations{}
		err := json.Unmarshal(dat, &locationsResp)
		if err != nil {
			return Locations{}, err
		}
		return locationsResp, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Locations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Locations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Locations{}, err
	}
	c.cache.Add(url, dat)
	locationsResp := Locations{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return Locations{}, err
	}

	return locationsResp, nil
}

type Locations struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
}
