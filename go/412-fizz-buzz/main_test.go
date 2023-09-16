package main

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	testcases := map[string]struct {
		input    int
		expected []string
	}{
		"case1": {input: 3, expected: []string{"1", "2", "Fizz"}},
		"case2": {input: 5, expected: []string{"1", "2", "Fizz", "4", "Buzz"}},
		"case3": {input: 15, expected: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}},
	}

	for caseName, tc := range testcases {
		t.Run(caseName, func(t *testing.T) {
			actual := fizzBuzz(tc.input)
			t.Logf("input: %v, expected: %v, actual: %v", tc.input, tc.expected, actual)

			for i, number := range actual {
				if number != tc.expected[i] {
					t.Errorf("Expected %v, got %v", tc.expected, actual)
				}
			}
		})
	}
}
