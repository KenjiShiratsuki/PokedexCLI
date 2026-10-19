package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config, args ...string) error {
	loc, err := cfg.pokeapiClient.ListLocations(cfg.NextLocations)
	if err != nil {
		return err
	}

	cfg.NextLocations = loc.Next
	cfg.PrevLocations = loc.Previous

	for _, area := range loc.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.PrevLocations == "" {
		return errors.New("You are already on the first page")
	}
	loc, err := cfg.pokeapiClient.ListLocations(cfg.PrevLocations)
	if err != nil {
		return err
	}

	cfg.NextLocations = loc.Next
	cfg.PrevLocations = loc.Previous
	for _, area := range loc.Results {
		fmt.Println(area.Name)
	}
	return nil
}
