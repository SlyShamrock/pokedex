package main
import (
	"fmt"
	"os"
	"github.com/slyshamrock/pokedex/internal/pokeapi"
	"github.com/slyshamrock/pokedex/internal/pokecache"
	"errors"	
	"math/rand"
)

type cliCommand struct {
	name string
	description string
	callback func(*locationConfig, []string) error
}

type locationConfig struct {
	cache *pokecache.Cache
	NextPage *string
	PrevPage *string
	pokedex map[string]pokeapi.Pokemon
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "Displays the next 20 locations on the map",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Displays the previous 20 locations on the map",
			callback: commandMapb,
		},
		"explore": {
			name: "explore",
			description: "Displays all Pokemon at map location",
			callback: commandExplore,
		},
		"catch": {
			name: "catch",
			description: "Attempts to catch a Pokemon",
			callback: commandCatch,
		},
		"inspect": {
			name: "inspect",
			description: "Shows information about captured Pokemon",
			callback: commandInspect,
		},
		"pokedex": {
			name: "pokedex",
			description: "Lists Pokemon that have been caught",
			callback: commandPokedex,
		},
	}
}

func commandExit(config *locationConfig, names []string) error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp(config *locationConfig, names []string) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")				
	cmdMap := getCommands()
	for _, cmd := range cmdMap {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(config *locationConfig, names []string) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if config.NextPage != nil {
		url = *config.NextPage
	}

	resp, err := pokeapi.GetLocation(url, config.cache)
	if err != nil {
		return err
	}
	
	for _, result := range resp.Results {
		fmt.Printf("%s\n", result.Name)		 	
	}

	config.NextPage = resp.Next
	config.PrevPage = resp.Previous

	return nil
}

func commandMapb(config *locationConfig, names []string) error {
	if config.PrevPage == nil {
		fmt.Printf("You're on the first page\n")
		return nil
	}
	url := *config.PrevPage
	resp, err := pokeapi.GetLocation(url, config.cache)
	if err != nil {
		return err
	}

	for _, result := range resp.Results {
		fmt.Printf("%s\n", result.Name)
	}

	config.NextPage = resp.Next
	config.PrevPage = resp.Previous			
	return nil
}

func commandExplore(config *locationConfig, names []string) error {
	if len(names) != 1 {
		return errors.New("you must provide a location")
	}
	
	locName := names[0]
	fmt.Printf("Exploring %s\n", locName)
	
	resp, err := pokeapi.GetLocationDetails(locName, config.cache)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, result := range resp.PokemonEncounters {
		fmt.Printf("- %s\n", result.Pokemon.Name)
	}

	return nil
}

func commandCatch(config *locationConfig, names []string) error {
	if len(names) != 1 {
		return errors.New("you must provide a Pokemon name")
	}

	pokeName := names[0]	
	fmt.Printf("Throwing a Pokeball at %s...\n", pokeName)

	resp, err := pokeapi.GetPokemon(pokeName, config.cache)
	if err != nil {
		return err
	}
	roll := rand.Intn(resp.BaseExperience + 150)
	if roll < 50 {
		fmt.Printf("%s was caught!\n", resp.Name)
		config.pokedex[resp.Name] = resp
	} else {
		fmt.Printf("%s escaped!\n", resp.Name)
	}

	return nil
}

func commandInspect(config *locationConfig, names []string) error {
	if len(names) != 1 {
		return errors.New("you must provide a Pokemon name")
	}

	pokeName := names[0]
	if pokemon, ok := config.pokedex[pokeName]; ok {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Println("Stats: ")
		for _, stat := range pokemon.Stats {
			fmt.Printf(" -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types: ")
		for _, tp := range pokemon.Types {
			fmt.Printf(" -%s\n", tp.Type.Name)
		}
	} else {
		fmt.Printf("Pokemon not in Pokedex")
	}
	return nil
}

func commandPokedex(config *locationConfig, names []string) error {
	fmt.Println("Your Pokedex:")
	for key, _ := range config.pokedex {
		fmt.Printf(" -%s\n", key)
	}
	return nil
}