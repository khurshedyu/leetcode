package main

import "testing"

func TestMaximumWealth(t *testing.T) {
	testcases := map[string]struct {
		input    [][]int
		expected int
	}{
		"case1": {input: [][]int{{1, 2, 3}, {3, 2, 1}}, expected: 6},
		"case2": {input: [][]int{{1, 5}, {7, 3}, {3, 5}}, expected: 10},
		"case3": {input: [][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}, expected: 17},
	}

	for caseName, tc := range testcases {
		t.Run(caseName, func(t *testing.T) {
			actual := maximumWealth(tc.input)
			t.Logf("input: %v, expected: %v, actual: %v", tc.input, tc.expected, actual)

			if actual != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, actual)
			}
		})
	}
}
