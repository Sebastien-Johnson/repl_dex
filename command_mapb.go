package main

import (
	"fmt"
	"log"
)

func commandMapb(cfg *config) error {
	//create new client from constructor via cfg
	//get resp and err from location area request
	// cfg.endpointIndex -= 20
	// endpoint := fmt.Sprintf("/location-area?offset=%v&limit=20", cfg.endpointIndex)
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreaUrl)
	//check for errors
	if err != nil {
		log.Fatal(err)
	}
	//print resp
	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" -%s\n", area.Name)
	}
	cfg.nextLocationAreaUrl = resp.Next
	cfg.prevLocationAreaUrl = resp.Previous
	return nil
}