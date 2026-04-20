package main

import (
	"fmt"
)

func commandMap(cfg *config, locationName *string) error {
	//create new client from constructor via cfg
	//get resp and err from location area 
	locationsResp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaUrl)
	//check for errors
	if err != nil {
		return err
	}

	//sets config properties to new response
	cfg.nextLocationAreaUrl = locationsResp.Next
	cfg.prevLocationAreaUrl = locationsResp.Previous

	//print resp
	fmt.Println("Location areas:")
	for _, area := range locationsResp.Results {
		fmt.Printf(" -%s\n", area.Name)
	}
	
	return nil
}
