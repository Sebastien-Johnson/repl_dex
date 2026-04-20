package main

import (
	"time"

	"github.com/sebastien-johnson/repl_dex/internal/pokeapi"
)




func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := config{
		pokeapiClient: pokeClient,
		pokemonCaught: make(map[string]pokeapi.PokemonDetailsResp),
	}
	startRepl(&cfg)
}

