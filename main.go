package main

import (
	"fmt"
 	"bufio"
	"os"
	"time"
	"github.com/slyshamrock/pokedex/internal/pokecache"
	"github.com/slyshamrock/pokedex/internal/pokeapi"	
)



func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := &locationConfig{
		cache: pokecache.NewCache(5 * time.Second),
		pokedex: make(map[string]pokeapi.Pokemon),
	}	
	for ;; {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		newText := scanner.Text()
		formattedInput := cleanInput(newText)
		if len(formattedInput) == 0 {
			continue
		}
		firstWord := formattedInput[0]					
		cmdMap := getCommands()
		value, ok := cmdMap[firstWord]		
		if !ok {
			fmt.Printf("Unknown command\n")
			continue
		}
		err := value.callback(cfg, formattedInput[1:])
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}				
	}
}

