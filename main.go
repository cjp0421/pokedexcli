package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex >")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		cleanedInput := cleanInput(string(input))
		if len(cleanedInput) > 0 {
			fmt.Printf("Your command was: %v\n", cleanedInput[0])
		}
	}
}
func cleanInput(text string) []string {
	// fmt.Println(text)
	lowerText := strings.ToLower(text)
	cutset := " "
	textStrings := (strings.Split(strings.Trim(lowerText, cutset), " "))
	return textStrings
}
