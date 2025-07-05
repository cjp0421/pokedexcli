package utilities

import "strings"

func CleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	cutset := " "
	textStrings := (strings.Split(strings.Trim(lowerText, cutset), " "))
	return textStrings
}
