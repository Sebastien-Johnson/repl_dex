package pokeapi

import (
	"fmt"
	"encoding/json"
	"net/http"
	"io"
)

func (c *Client) ListAreaDetails(areaName *string) (areaDetails, error) {
	fullUrl := baseUrl + *areaName

	if dat, ok := c.cache.Get(fullUrl) ; ok{
		fmt.Println("cache hit!")
		//create empty var for json response
		areaDetailsResp := areaDetails{}
		//check json data for error and/or write to pointer to var
		err := json.Unmarshal(dat, &areaDetailsResp)
		if err != nil {
			return areaDetails{}, err
		}
		//returns valid location area and nil error early
		return areaDetailsResp, nil
	}
	fmt.Println("cache miss!")

	//Get new request, preps for sending
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil{
		//return empty location area & error
		return areaDetails{}, err
	}
	//now request is verfied/valid

	//.Do(req) executes request, and gets http resp
	resp, err := c.httpClient.Do(req)
	if err != nil{
		return areaDetails{}, err
	}

	//close resp body obj by default
	defer resp.Body.Close()

	//read through json data
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return areaDetails{}, err
	}
	
	//create empty var for json response
	areaDetailsResp := areaDetails{}
	//check json data for error and/or write to pointer to var
	err = json.Unmarshal(data, &areaDetailsResp)
	if err != nil {
		return areaDetails{}, err
	}
	//adds new valid fullUrl to cache
	c.cache.Add(fullUrl, data)
	//returns valid location area and nil error
	return areaDetailsResp, nil
}