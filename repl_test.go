package main

import (
	"testing"

	"github.com/cjp0421/pokedexcli/utilities"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " a string of pokemon creatures ",
			expected: []string{"a", "string", "of", "pokemon", "creatures"},
		},
	}

	for _, c := range cases {
		actual := utilities.CleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Length of actual does not match expected")
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if expectedWord != word {
				t.Errorf("Expected word does not match word")
			}
		}
	}
}
