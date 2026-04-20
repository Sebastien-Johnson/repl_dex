package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/sebastien-johnson/repl_dex/internal/pokeapi"
)

type config struct {
	//allows efficient re-use of client
	pokeapiClient 		pokeapi.Client
	nextLocationAreaUrl *string
	prevLocationAreaUrl *string
	//pointer allows it to be 'nil' when non exist
}

//input and command data, passes const config
func startRepl(cfg *config) {
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
		areaName := ""
		if len(input) > 1 {
			areaName = input[1]
		}
		
		//access command in command map via name-index
		command, exists := getCommands()[commandName]

		if exists { // if exists and checks for error
			err := command.callback(cfg, &areaName)
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


// form struct for command data, set callbacks to accept config
type cliCommand struct {
	name string
	description string
	callback func(*config, *string) error 
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
			callback: commandMap, //callback
		},
		"mapb": {
			name: "mapb",
			description: "Displays previous 20 area names",
			callback: commandMapb, //callback
		},
		"explore": {
			name: "explore",
			description: "Shows pokemon in a given area",
			callback: commandExplore,
		},
	}
}
	

