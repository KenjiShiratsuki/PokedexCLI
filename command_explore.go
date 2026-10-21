package main

import (
	"errors"
	"fmt"
)

func explore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("Please enter an area name to explore it.")
	}
	locationInfo, err := cfg.pokeapiClient.GetEncounters(args[0])
	if err != nil {
		return err
	}
	for _, pokemon := range locationInfo.Encounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}
