package pokeapi

import (
	"fmt"
	"net/http"
	"io"
	"encoding/json"
)

func GetLocation(url string) (LocationAreasResp, error) {
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
		return locations, nil
	}