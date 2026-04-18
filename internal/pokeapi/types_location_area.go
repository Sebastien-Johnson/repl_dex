package pokeapi

//location area response struct from jlint + json to go
type LocationAreasResp struct {
	Count    int    `json:"count"`
	Next     *string   `json:"next"` //pointer to url
	Previous *string    `json:"previous"` //pointer to url
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}