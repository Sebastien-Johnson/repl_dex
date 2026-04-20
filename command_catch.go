package main

import (
	"fmt"
	"math/rand"
	"strings"
)


func commandCatch(cfg *config, pokemonName *string) error {
	endpoint := "pokemon/"+*pokemonName

	pokemonResp, err := cfg.pokeapiClient.ListPokemonDetails(&endpoint)

	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", *pokemonName)

	randRange := rand.Intn(700)

	catchRate := pokemonResp.BaseExperience
	if randRange > catchRate {
		fmt.Printf("%s was caught!\n\n", strings.Title(*pokemonName))
		fmt.Printf("You may now inspect it with the inspect command.")
		cfg.pokemonCaught[*pokemonName] = pokemonResp
	} else {
		fmt.Printf("Failed to catch %s );\n\n", strings.Title(*pokemonName))
	}
	

	
	fmt.Print("\n")
	return nil
}