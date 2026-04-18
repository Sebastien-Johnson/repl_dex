package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//input and command data
func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println(("Pokedex > "))
		scanner.Scan()

		input := cleanInput(scanner.Text())
		//check for empty input
		if len(input) == 0 {
			continue
		}
		//grab first word from cleaned input
		commandName := input[0]
		//access command in command map via name-index
		command, exists := getCommands()[commandName]
		if exists { // if exists and checks for error
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else { //if doesnt exist
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}


// form struct for command data
type cliCommand struct {
	name string
	description string
	callback func() error 
}

//function that returns map with string:struct pairs that contains callback to funcs, accepts string and returns cli command structs
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit, //callback
		},
		"help": {
			name: "help",
			description: "Display a help message",
			callback: commandHelp, //callback
		},
		"map": {
			name: "map",
			description: "List location areas",
			callback: commandMap,
		},
		// "mapb": {
		// 	name: "mapb",
		// 	description: "Displays previous 20 area names",
		// 	callback: commandMapb,
		//},
	}
}
	

