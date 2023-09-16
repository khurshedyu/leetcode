package main

import "testing"

func TestNumberOfSteps(t *testing.T) {
	testcases := map[string]struct {
		input    int
		expected int
	}{
		"case1": {input: 14, expected: 6},
		"case2": {input: 8, expected: 4},
		"case3": {input: 123, expected: 12},
	}

	for caseName, tc := range testcases {
		t.Run(caseName, func(t *testing.T) {
			actual := numberOfSteps(tc.input)
			t.Logf("input: %v, expected: %v, actual: %v", tc.input, tc.expected, actual)

			if actual != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, actual)
			}
		})
	}
}
