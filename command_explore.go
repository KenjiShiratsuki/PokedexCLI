package main

import (
	"errors"
	"fmt"
)

func explore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("Please enter an area name to explore it.")
	}
	pokemonList, err := cfg.pokeapiClient.GetEncounters(args[0])
	if err != nil {
		return err
	}
	for _, pokemon := range pokemonList.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}
