package pokeapi

import (
	"encoding/json"
	"net/http"
	"io"
	"fmt"
)

// takes client, returns location area response struct and error
//c == client struct
func (c *Client, ) ListLocationAreas(pageUrl *string) (LocationAreasResp, error) {
	//combine endpoint with const base url
	endpoint := "/location-area"
	fullUrl := baseUrl+endpoint

	if pageUrl != nil {
		fullUrl = *pageUrl
	} 

	//check the cache
	dat, ok := c.cache.Get(fullUrl)
	if ok {
		//cache hits
		//create empty var for json response
		fmt.Println("cache hit!")
		locationAreasResp := LocationAreasResp{}
		//check json data for error and/or write to pointer to var
		err1 := json.Unmarshal(dat, &locationAreasResp)
		if err1 != nil {
			return LocationAreasResp{}, err1
		}
		//returns valid location area and nil error early
		return locationAreasResp, nil
	}
	fmt.Println("cache miss!")

	// get request with NewRequest, no body
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil{
		//return empty location area & error
		return LocationAreasResp{}, err
	}
	//now request is verfied/valid

	//.Do(req) executes request, check for errors
	resp, err := c.httpClient.Do(req)
	if err != nil{
		return LocationAreasResp{}, err
	}
	//close resp body obj
	defer resp.Body.Close()

	//check status code
	if resp.StatusCode >299 {
		return LocationAreasResp{}, fmt.Errorf("Response failed with status code: %d", resp.StatusCode)
	}

	//read through json data
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	//create empty var for json response
	locationAreasResp := LocationAreasResp{}
	//check json data for error and/or write to pointer to var
	err2 := json.Unmarshal(data, &locationAreasResp)
	if err2 != nil {
		return LocationAreasResp{}, err2
	}
	//adds new valid fullUrl to cache
	c.cache.Add(fullUrl, dat)
	//returns valid location area and nil error
	return locationAreasResp, nil
}