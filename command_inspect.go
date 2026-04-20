package main

import (
	"errors"
	"fmt"
	"strings"
)

func commandInspect(cfg *config, pokemonName *string) error {
	pokemonStats, ok := cfg.pokemonCaught[*pokemonName]

	if !ok {
		return errors.New(fmt.Sprintf("%s has not been caught yet", strings.Title(*pokemonName)))
	}

	fmt.Printf("Name: %s\n", pokemonStats.Name)
	fmt.Printf("Height: %s\n", pokemonStats.Height)
	fmt.Printf("Weight: %s\n", pokemonStats.Weight)
	fmt.Print("Stats: \n")
	for _, stat := range pokemonStats.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Print("Types: \n")
	for _, typ := range pokemonStats.Types {
		fmt.Printf("	- %s\n", typ.Type.Name)
	}

	return nil
}