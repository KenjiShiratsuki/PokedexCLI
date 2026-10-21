package main

import (
	"time"

	"github.com/KenjiShiratsuki/PokedexCLI/internal/pokeapi"
)

func main() {
	cfg := &config{
		commandList:   getCommands(),
		pokeapiClient: pokeapi.NewClient(5 * time.Second),
		Pokedex:       make(map[string]pokeapi.Pokemon),
	}
	startRepl(cfg)
}
