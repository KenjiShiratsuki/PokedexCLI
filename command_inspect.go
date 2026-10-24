package main

import (
	"errors"
	"fmt"
)

func inspect(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("Please enter a Pokemon name to inspect it.")
	}
	if pkmn, exists := cfg.Pokedex[args[0]]; exists {
		fmt.Printf("Name: %v\n", pkmn.Name)
		fmt.Printf("Height: %v\n", pkmn.Height)
		fmt.Printf("Weight: %v\n", pkmn.Weight)
		fmt.Printf("Stats:\n")
		for _, stat := range pkmn.Stats {
			fmt.Printf(" - %s: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Printf("Types:\n")
		for _, t := range pkmn.Types {
			fmt.Printf(" - %s\n", t.Type.Name)
		}
	} else {
		return errors.New("You haven't caught that Pokemon yet!")
	}
	return nil
}
