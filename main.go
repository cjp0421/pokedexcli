package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
	// cleanInput("hello world")
}
func cleanInput(text string) []string {
	fmt.Println(text)
	lowerText := strings.ToLower(text)
	cutset := " "
	textStrings := (strings.Split(strings.Trim(lowerText, cutset), " "))
	return textStrings
}
