package main

import (
	"time"

	"github.com/sebastien-johnson/repl_dex/internal/pokeapi"
)

type config struct {
	//allows efficient re-use of client
	pokeapiClient 		pokeapi.Client
	nextLocationAreaUrl *string
	prevLocationAreaUrl *string
	endpointIndex 		int
	//pointer allows it to be 'nil' when non exist
}


func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
	}
	startRepl(&cfg)
}

