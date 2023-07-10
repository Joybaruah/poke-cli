package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	endpoint := "/location-area"
	fullURL := baseUrl + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}
	// Check if cache exist

	data, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("cache data")
		locationAreasResp := LocationAreasResp{}
		err := json.Unmarshal(data, &locationAreasResp)
		if err != nil {
			return LocationAreasResp{}, err
		}
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("server error code: %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	locationAreasResp := LocationAreasResp{}
	err = json.Unmarshal(data, &locationAreasResp)
	if err != nil {
		return LocationAreasResp{}, err
	}

	c.cache.Add(fullURL, data)

	return locationAreasResp, nil
}

func (c *Client) GetLocation(locationArea *string) (LocationArea, error) {
	endpoint := "/location-area/" + *locationArea
	fullURL := baseUrl + endpoint

	data, ok := c.cache.Get(fullURL)
	if ok {
		locationAreaResp := LocationArea{}
		err := json.Unmarshal(data, &locationAreaResp)
		if err != nil {
			return LocationArea{}, err
		}
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("server error code: %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationAreaResp := LocationArea{}
	err = json.Unmarshal(data, &locationAreaResp)
	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(fullURL, data)

	return locationAreaResp, nil
}
