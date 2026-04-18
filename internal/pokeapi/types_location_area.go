package pokeapi

//location area response struct from jlint + json to go
type LocationAreasResp struct {
	Count    int    `json:"count"`
	Next     *string   `json:"next"`
	Previous *string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}