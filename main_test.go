package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	testcases := []struct {
		input 	 string
		expected []string
	}{
		{
			input: 	  "  hello  world  ",
			expected: []string{ "hello", "world" },
		},
		{
			input:	  "Charmander Bulbasaur PIKACHU",
			expected: []string{ "charmander", "bulbasaur", "pikachu" },
		},
	}

	for _, tc := range testcases {
		actual := cleanInput(tc.input)
		if len(actual) != len(tc.expected) {
			t.Errorf("Error -> Length of actual words doesn't match length of expected words\nExpected: %v\n  Actual: %v", tc.expected, actual)
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedword := tc.expected[i]
			if word != expectedword {
				t.Errorf("Error -> Actual words doesn't match expected words\nExpected: %v\n  Actual: %v", tc.expected, actual)
				break
			}
		}
	}
}
