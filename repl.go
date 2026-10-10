package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(state *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}
		command, exists := state.commandList[words[0]]
		if exists {
			command.callback(state)
		} else {
			fmt.Println("Unknown Command")
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
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

func commandExit(cfg *config) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return fmt.Errorf("Error exiting program, use ctrl+c to terminate cli")
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	words := strings.Fields(lowered)
	return words
}

type config struct {
	commandList map[string]cliCommand
}
