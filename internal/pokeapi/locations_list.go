package pokeapi

import (
	"fmt"
	"net/http"
	"io"
	"encoding/json"
	"github.com/slyshamrock/pokedex/internal/pokecache"	
)

func GetLocation(url string, cache *pokecache.Cache) (LocationAreasResp, error) {
	if data, ok := cache.Get(url); ok {
		var locations LocationAreasResp
		if err := json.Unmarshal(data, &locations); err != nil {
			return LocationAreasResp{}, fmt.Errorf("error creating request: %w", err)
		}
		return locations, nil
	}
	
	res, err := http.Get(url)
	if err != nil {
		return LocationAreasResp{}, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)	
	if err != nil {
		return LocationAreasResp{}, fmt.Errorf("error reading response: %w", err)
	}
	if res.StatusCode > 299 {
		return LocationAreasResp{}, fmt.Errorf("bad status: %d", res.StatusCode)
	}

	var locations LocationAreasResp
	if err := json.Unmarshal(data, &locations); err != nil {
		return LocationAreasResp{}, fmt.Errorf("error decoding response: %w", err)
	}
	cache.Add(url, data)
	return locations, nil
}

func GetLocationDetails(name string, cache *pokecache.Cache) (AreaDetails, error) {
	fullUrl := "https://pokeapi.co/api/v2/location-area/" + name
	if data, ok := cache.Get(fullUrl); ok {
		var locations AreaDetails
		if err := json.Unmarshal(data, &locations); err != nil {
			return AreaDetails{}, fmt.Errorf("error creating request: %w", err)
		}
		return locations, nil
	}
	
	res, err := http.Get(fullUrl)
	if err != nil {
		return AreaDetails{}, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)	
	if err != nil {
		return AreaDetails{}, fmt.Errorf("error reading response: %w", err)
	}
	if res.StatusCode > 299 {
		return AreaDetails{}, fmt.Errorf("bad status: %d", res.StatusCode)
	}

	var locations AreaDetails
	if err := json.Unmarshal(data, &locations); err != nil {
		return AreaDetails{}, fmt.Errorf("error decoding response: %w", err)
	}
	cache.Add(fullUrl, data)
	return locations, nil
}