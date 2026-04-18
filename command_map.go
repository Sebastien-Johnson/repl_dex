package main

import (
	"fmt"
	"log"
	"github.com/sebastien-johnson/repl_dex/internal/pokeapi"
)

func commandMap() error {
	//create new client from constructor
	pokeapiClient := pokeapi.NewClient()
	//get resp and err from location area request
	resp, err := pokeapiClient.ListLocationAreas()
	//check for errors
	if err != nil {
		log.Fatal(err)
	}
	//print resp
	fmt.Println(resp)
	return nil
}

// func commandMapb() error {
// 	pokeapiClient := pokeapi.NewClient()
// 	//get resp
// 	resp, err := pokeapiClient.ListLocationAreas()
// 	//check for errors
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	//print resp
// 	fmt.Println("Location areas:")
// 	for _, area := range resp.Results{
// 		fmt.Printf("-%s\n", area.Name)
// 	}
// 	return nil
// }