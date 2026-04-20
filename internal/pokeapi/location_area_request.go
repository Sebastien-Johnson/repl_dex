package pokeapi

import (
	"encoding/json"
	"net/http"
	"io"
	
)

// takes client, returns location area response struct and error
//c == client struct
func (c *Client) ListLocationAreas(pageUrl *string) (LocationAreasResp, error) {
	//combine endpoint with const base url
	endpoint := "location-area"
	fullUrl := baseUrl+endpoint
	

	if pageUrl != nil {
		fullUrl = *pageUrl
	} 

	//check if fullUrl data is in cache
	if dat, ok := c.cache.Get(fullUrl) ; ok{
		// fmt.Println("cache hit!")
		//create empty var for json response
		locationAreasResp := LocationAreasResp{}
		//check json data for error and/or write to pointer to var
		err := json.Unmarshal(dat, &locationAreasResp)
		if err != nil {
			return LocationAreasResp{}, err
		}
		//returns valid location area and nil error early
		return locationAreasResp, nil
	}
	// fmt.Println("cache miss!")

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

	//read through json data
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	//create empty var for json response
	locationAreasResp := LocationAreasResp{}
	//check json data for error and/or write to pointer to var
	err = json.Unmarshal(data, &locationAreasResp)
	if err != nil {
		return LocationAreasResp{}, err
	}
	//adds new valid fullUrl to cache
	c.cache.Add(fullUrl, data)
	//returns valid location area and nil error
	return locationAreasResp, nil
}