package pokeapi

import (
	"net/http"
	"time"

	"github.com/sebastien-johnson/repl_dex/internal/pokecache"
)

//create const for base url
const baseUrl = "https://pokeapi.co/api/v2"

//Create own client struct
type Client struct {
	//store cache in client
	cache pokecache.Cache
	httpClient http.Client
}

//create constructor func to generate and return new clients
func NewClient(cacheInterval time.Duration) Client {
	return Client {
		cache: pokecache.NewCache(cacheInterval),
		httpClient : http.Client{
			// add time out to client
			Timeout : time.Minute,
		},
	}
}