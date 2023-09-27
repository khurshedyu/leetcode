package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	testcases := map[string]struct {
		num      int
		expected bool
	}{
		"case1": {num: 121, expected: true},
		"case2": {num: -121, expected: false},
		"case3": {num: 10, expected: false},
	}

	for caseName, tc := range testcases {
		t.Run(caseName, func(t *testing.T) {
			actual := isPalindrome(tc.num)
			t.Logf("num: %v, expected: %v, actual: %v", tc.num, tc.expected, actual)

			if actual != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
