package pokeapi

import (
	"fmt"
	"net/http"
	"io"
	"encoding/json"
	"github.com/slyshamrock/pokedex/internal/pokecache"	
)

func GetPokemon(name string, cache *pokecache.Cache) (Pokemon, error) {	
	fullUrl := "https://pokeapi.co/api/v2/pokemon/" + name
	if data, ok := cache.Get(fullUrl); ok {
		var poke Pokemon
		if err := json.Unmarshal(data, &poke); err != nil {
			return Pokemon{}, fmt.Errorf("error creating request: %w", err)
		}
		return poke, nil
	}
	
	res, err := http.Get(fullUrl)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)	
	if err != nil {
		return Pokemon{}, fmt.Errorf("error reading response: %w", err)
	}
	if res.StatusCode > 299 {
		return Pokemon{}, fmt.Errorf("bad status: %d", res.StatusCode)
	}

	var poke Pokemon
	if err := json.Unmarshal(data, &poke); err != nil {
		return Pokemon{}, fmt.Errorf("error decoding response: %w", err)
	}
	cache.Add(fullUrl, data)
	return poke, nil
}
