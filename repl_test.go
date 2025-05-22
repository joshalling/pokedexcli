package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " Hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "gOlAnG is BaD",
			expected: []string{"golang", "is", "bad"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("Lengths differ. Expected %v, got %v", len(c.expected), len(actual))
			continue
		}

		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("Values differ.Expected %v, got %v", c.expected, actual)
			}
		}
	}
}
