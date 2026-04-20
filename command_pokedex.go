package main

import (
	"fmt"
)

func commandPokedex(cfg *config, pokedex *string) error {
	if len(cfg.pokemonCaught) != 0 {
		fmt.Print("Your Pokedex:\n")
		for _, mon := range cfg.pokemonCaught {
			fmt.Printf("-%s \n", mon.Name)
		}
	} else {
		fmt.Print("No pokemon caught yet!")
	}
	return nil
}