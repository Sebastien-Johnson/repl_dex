package main

import (
	"github.com/sebastien-johnson/repl_dex/internal/pokeapi"
	"log"
	"fmt"
)

func main() {
	//startRepl()
	pokeapiClient := pokeapi.NewClient()
	//get resp
	resp, err := pokeapiClient.ListLocationAreas()
	//check for errors
	if err != nil {
		log.Fatal(err)
	}
	//print resp
	fmt.Println(resp)
}

