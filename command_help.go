package main

import (
	"fmt"
)


func commandHelp(cfg *config, locationName *string) error {
	fmt.Println("")
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	//iterates through each command struct in map
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	} 
	fmt.Println("")
	return nil
}