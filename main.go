package main
import "fmt"
import "bufio"
import "os"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for ;; {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		newText := scanner.Text()
		formattedInput := cleanInput(newText)
		if len(formattedInput) == 0 {
			continue
		}
		firstWord := formattedInput[0]					
		commandMap := getCommands()
		value, ok := commandMap[firstWord]		
		if !ok {
			fmt.Printf("Unknown command\n")
			continue
		}
		err := value.callback()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}				
	}
}

