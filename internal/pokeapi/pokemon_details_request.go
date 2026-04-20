package pokeapi

import (
	"encoding/json"
	"net/http"
	"io"
	
)

func (c *Client) ListPokemonDetails(pageUrl *string) (PokemonDetailsResp, error) {
	fullUrl := baseUrl + *pageUrl

	if dat, ok := c.cache.Get(fullUrl) ; ok{
		// fmt.Println("cache hit!")
		//create empty var for json response
		pokemonDetailsResp := PokemonDetailsResp{}
		//check json data for error and/or write to pointer to var
		err := json.Unmarshal(dat, &pokemonDetailsResp)
		if err != nil {
			return PokemonDetailsResp{}, err
		}
		//returns valid location area and nil error early
		return pokemonDetailsResp, nil
	}
	// fmt.Println("cache miss!")

	//Get new request, preps for sending
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil{
		//return empty location area & error
		return PokemonDetailsResp{}, err
	}
	//now request is verfied/valid

	//.Do(req) executes request, and gets http resp
	resp, err := c.httpClient.Do(req)
	if err != nil{
		return PokemonDetailsResp{}, err
	}

	//close resp body obj by default
	defer resp.Body.Close()

	//read through json data
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonDetailsResp{}, err
	}
	
	//create empty var for json response
	pokemonDetailsResp := PokemonDetailsResp{}
	//check json data for error and/or write to pointer to var
	err = json.Unmarshal(data, &pokemonDetailsResp)
	if err != nil {
		return PokemonDetailsResp{}, err
	}
	//adds new valid fullUrl to cache
	c.cache.Add(fullUrl, data)
	//returns valid location area and nil error
	return pokemonDetailsResp, nil
}
