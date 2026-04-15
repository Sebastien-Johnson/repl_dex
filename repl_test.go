package main

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
// }