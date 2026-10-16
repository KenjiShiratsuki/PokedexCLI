package main

import "fmt"

func help(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n")
	allCommands := ""
	for _, command := range getCommands() {
		allCommands += command.name + ": " + command.description + "\n"
	}
	print(allCommands)
	return nil
}
