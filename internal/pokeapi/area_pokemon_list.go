package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) GetEncounters(areaName string) (LocationInfo, error) {
	url := baseURL + "/location-area/" + areaName
	if dat, ok := c.cache.Get(url); ok {
		pokemonResp := LocationInfo{}
		err := json.Unmarshal(dat, &pokemonResp)
		if err != nil {
			return LocationInfo{}, err
		}
		return pokemonResp, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationInfo{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationInfo{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 404 {
		return LocationInfo{}, errors.New("unknown location, please enter a known location within the pokeverse")
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationInfo{}, err
	}
	c.cache.Add(url, dat)
	pokemonResp := LocationInfo{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return LocationInfo{}, err
	}

	return pokemonResp, nil
}

type LocationInfo struct {
	Encounters   PokemonEncounters `json:"pokemon_encounters"`
	LocationArea NamedAPIResource  `json:"location"`
}

type PokemonEncounters []struct {
	Pokemon NamedAPIResource `json:"pokemon"`
}
type NamedAPIResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
