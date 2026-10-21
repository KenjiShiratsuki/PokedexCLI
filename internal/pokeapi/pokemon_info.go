package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemon string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemon
	if dat, ok := c.cache.Get(url); ok {
		pokemonResp := Pokemon{}
		err := json.Unmarshal(dat, &pokemonResp)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemonResp, nil
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
	if resp.StatusCode == 404 {
		return Pokemon{}, errors.New("unknown Pokemon, please enter a known Pokemon name")
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}
	c.cache.Add(url, dat)
	pokemonResp := Pokemon{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemonResp, nil
}

type Pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	IsDefault      bool   `json:"is_default"`
	Order          int    `json:"order"`
	Weight         int    `json:"weight"`
	Abilities      []struct {
		IsHidden bool             `json:"is_hidden"`
		Slot     int              `json:"slot"`
		Ability  NamedAPIResource `json:"ability"`
	} `json:"abilities"`
	Forms     []NamedAPIResource `json:"forms"`
	HeldItems []struct {
		Item           NamedAPIResource `json:"item"`
		VersionDetails []struct {
			Rarity  int              `json:"rarity"`
			Version NamedAPIResource `json:"version"`
		} `json:"version_details"`
	} `json:"held_items"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Moves                  []struct {
		Move                NamedAPIResource `json:"move"`
		VersionGroupDetails []struct {
			LevelLearnedAt  int              `json:"level_learned_at"`
			VersionGroup    NamedAPIResource `json:"version_group"`
			MoveLearnMethod NamedAPIResource `json:"move_learn_method"`
			Order           int              `json:"order"`
		} `json:"version_group_details"`
	} `json:"moves"`
	Species NamedAPIResource `json:"species"`
	Stats   []struct {
		BaseStat int              `json:"base_stat"`
		Effort   int              `json:"effort"`
		Stat     NamedAPIResource `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int              `json:"slot"`
		Type NamedAPIResource `json:"type"`
	} `json:"types"`
}
