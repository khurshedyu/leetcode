package main

import "testing"

func TestRunningSum(t *testing.T) {
	testcases := map[string]struct {
		input    []int
		expected []int
	}{
		"case1": {input: []int{1, 2, 3, 4}, expected: []int{1, 3, 6, 10}},
		"case2": {input: []int{1, 1, 1, 1, 1}, expected: []int{1, 2, 3, 4, 5}},
		"case3": {input: []int{3, 1, 2, 10, 1}, expected: []int{3, 4, 6, 16, 17}},
		"case4": {input: []int{0, 0, 0, 0}, expected: []int{0, 0, 0, 0}},
	}

	for caseName, tc := range testcases {
		t.Run(caseName, func(t *testing.T) {
			actual := runningSum(tc.input)
			t.Logf("input: %v, expected: %v, actual: %v", tc.input, tc.expected, actual)

			for i, number := range actual {
				if number != tc.expected[i] {
					t.Errorf("Expected %d, got %d", tc.expected, actual)
				}
			}
		})
	}
}
