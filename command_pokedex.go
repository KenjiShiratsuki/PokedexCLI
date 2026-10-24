package main

import (
	"fmt"
)

func pokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for _, pkmn := range cfg.Pokedex {
		fmt.Printf(" - %s\n", pkmn.Name)
	}
	return nil
}
