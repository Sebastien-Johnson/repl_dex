package pokeapi

import (
	"net/http"
	"time"
)

//create const for base url
const baseUrl = "https://pokeapi.co/api/v2"

//Create own client struct
type Client struct {
	httpClient http.Client
}

//create constructor func to generate and return new clients
func NewClient() Client {
	return Client {
		httpClient : http.Client{
			// add time out to client
			Timeout : time.Minute,
		},
	}
}