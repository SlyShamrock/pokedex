package main
import (
	"fmt"
	"os"
	"github.com/slyshamrock/pokedex/internal/pokeapi"
)

type cliCommand struct {
	name string
	description string
	callback func(*locationConfig) error
}

type locationConfig struct {
	NextPage *string
	PrevPage *string
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
	}
}

func commandExit(config *locationConfig) error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp(config *locationConfig) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")				
	cmdMap := getCommands()
	for _, cmd := range cmdMap {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(config *locationConfig) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if config.NextPage != nil {
		url = *config.NextPage
	}

	resp, err := pokeapi.GetLocation(url)
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

func commandMapb(config *locationConfig) error {
	if config.PrevPage == nil {
		fmt.Printf("You're on the first page\n")
		return nil
	}
	url := *config.PrevPage
	resp, err := pokeapi.GetLocation(url)
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