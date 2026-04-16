package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input  string
		expected []string
	}{
		{//test case
			input: " hello world ",
			expected: []string{"hello", "wolrd"},
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
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("lengths dont match: %v vs %v", actual, c.expected)
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("cleanInput(%v) == %v, expected %v", c.input, actual, c.expected)
			}
		}
	}
}
// func TestSplit(t *testing.T) {
//     tests := map[string]struct {
//         input string
//         sep   string
//         want  []string
//     }{
//         "simple":       {input: "a/b/c", sep: "/", want: []string{"a", "b", "c"}},
//         "wrong sep":    {input: "a/b/c", sep: ",", want: []string{"a/b/c"}},
//         "no sep":       {input: "abc", sep: "/", want: []string{"abc"}},
//         "trailing sep": {input: "a/b/c/", sep: "/", want: []string{"a", "b", "c"}},
//     }

//     for name, tc := range tests {
//         t.Run(name, func(t *testing.T) {
//             got := Split(tc.input, tc.sep)
//             diff := cmp.Diff(tc.want, got)
//             if diff != "" {
//                 t.Fatalf(diff)
//             }
//         })
//     }
