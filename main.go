package main

import "fmt"
import "bufio"
import "os"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := &locationConfig{}
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
		err := value.callback(cfg)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}				
	}
}

