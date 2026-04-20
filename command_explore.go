package main

import (
	"fmt"
)


func commandExplore(cfg *config, locationName *string) error {
	url := "location-area/"
	fullUrl := url+*locationName
	
	detailsResp, err := cfg.pokeapiClient.ListAreaDetails(&fullUrl)

	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", *locationName)
	fmt.Print("Found Pokemon:\n")
	for _, mon := range detailsResp.PokemonEncounters {
		fmt.Printf(" -%s\n", mon.Pokemon.Name)
	}
	
	return nil
}