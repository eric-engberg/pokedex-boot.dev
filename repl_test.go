package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	tests := []struct {
		input string
		expected []string
	}{
		{input: "help", expected: []string{"help"}},
		{input: "help me", expected: []string{"help", "me"}},
		{input: "help me please", expected: []string{"help", "me", "please"}},
		{input: "help ME please", expected: []string{"help", "me", "please"}},
	}

	for _, test := range tests {
		result := cleanInput(test.input)

		if len(result) != len(test.expected) {
			t.Errorf("Expected %d words, got %d", len(test.expected), len(result))
		}

		for i := range result {
			word := result[i]
			expectedWord := test.expected[i]
			if word != expectedWord {
				t.Errorf("Expected %s, got %s", expectedWord, word)
			}
		}
	}
}

func TestCleanInput_Empty(t *testing.T) {
	result := cleanInput("")
	if len(result) != 0 {
		t.Errorf("Expected empty slice, got %v", result)
	}
}
