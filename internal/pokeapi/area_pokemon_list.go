package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetEncounters(areaName string) (PokemonList, error) {
	url := baseURL + "/location-area/" + areaName
	if dat, ok := c.cache.Get(url); ok {
		pokemonResp := PokemonList{}
		err := json.Unmarshal(dat, &pokemonResp)
		if err != nil {
			return PokemonList{}, err
		}
		return pokemonResp, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonList{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonList{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonList{}, err
	}
	c.cache.Add(url, dat)
	pokemonResp := PokemonList{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return PokemonList{}, err
	}

	return pokemonResp, nil
}

type PokemonList struct {
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}
