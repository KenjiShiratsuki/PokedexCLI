package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}
		found := false
		for _, command := range commands {
			if command.name == words[0] {
				found = true
				command.callback()
			}
		}
		if found == false {
			fmt.Println("Unknown command")
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	commands := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    help,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
	return commands
}

func help() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n")
	allCommands := ""
	for _, command := range getCommands() {
		allCommands += command.name + ": " + command.description + "\n"
	}
	print(allCommands)
	return nil
}

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return fmt.Errorf("Error exiting program, use ctrl+c to terminate cli")
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	words := strings.Fields(lowered)
	return words
}
