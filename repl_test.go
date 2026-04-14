package 

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	casses := []struct {
		input  string
		expected []string
	}{
		{//test case
			input: " hello world ",
			expected []string{"hello", "wolrd"},
		},
		//add more cases
	}
	//loop over cases and run them all
	for _, c := range cases {
		actual := cleanInput(c.input)
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
		}
	}
}
