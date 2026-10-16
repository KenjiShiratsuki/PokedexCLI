package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return fmt.Errorf("Error exiting program, use ctrl+c to terminate cli")
}
