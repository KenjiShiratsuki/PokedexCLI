package main

import (
	"errors"
	"fmt"
	"math/rand"
)

const catchScale = 1000
const catchModifier = 0.9
const minCatchChance = 50

func catch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("please enter a pokemon name to try and catch it.")
	}
	pokemon := args[0]
	pokeInfo, err := cfg.pokeapiClient.GetPokemon(pokemon)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
	baseChance := catchScale - pokeInfo.BaseExperience
	catchChance := max(int(float64(baseChance)*catchModifier), minCatchChance)
	caught := rand.Intn(catchScale) < catchChance
	if caught {
		fmt.Printf("You caught a %s!\n", pokemon)
		cfg.Pokedex[pokemon] = pokeInfo
		return nil
	}
	fmt.Printf("%s broke free of the Pokeball!\n", pokemon)
	return nil
}
