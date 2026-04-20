package main

import (
	"fmt"
	"errors"
)

func commandMapb(cfg *config, locationName *string) error {
	if cfg.prevLocationAreaUrl == nil {
		return errors.New("You're on the first page")
	}
	//create new client from constructor via cfg
	//get resp and err from location area request
	locationResp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreaUrl)
	//check for errors
	if err != nil {
		return err
	}
	
	//sets config properties to new response
	cfg.nextLocationAreaUrl = locationResp.Next
	cfg.prevLocationAreaUrl = locationResp.Previous

	//print resp
	fmt.Println("Location areas:")
	for _, area := range locationResp.Results {
		fmt.Printf(" -%s\n", area.Name)
	}
	
	return nil
}