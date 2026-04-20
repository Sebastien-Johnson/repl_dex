package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	//create testing struct
	cases := []struct {
		input  string
		expected []string
	}{
		{//test case
			input: " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input: "does this even work?",
			expected: []string{"does", "this", "even", "work?"},
		},
		{
			input: "extra white space ",
			expected: []string{"extra", "white", "space"},
		},
		//add more cases
	}
	//loop over cases and run them all
	for _, c := range cases {
		//clean the input
		actual := cleanInput(c.input)
		//verify by length
		if len(actual) != len(c.expected) {
			t.Errorf("lengths dont match: %v vs %v", actual, c.expected)
			continue
		}
		//compare each word in input to expected
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("cleanInput(%v) == %v, expected %v", c.input, actual, c.expected)
			}
		}
	}
}

