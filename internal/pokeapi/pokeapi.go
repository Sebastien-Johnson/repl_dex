package pokeapi

import (
	"net/http"
	"time"
)

const baseUrl = "https://pokeapi.co/api/v2"

//Create own client to include time out
type Client struct {
	httpClient http.Client
}

//create constructor func to generate new clients
func NewClient() Client {
	return Client {
		httpClient : http.Client{
		Timeout : time.Minute,
		},
	}
}