package main

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	testcases := map[string]struct {
		str      string
		expected int
	}{
		"case1": {str: "abcabcbb", expected: 3},
		"case2": {str: "bbbbbb", expected: 1},
		"case3": {str: "pwwkew", expected: 3},
	}

	for caseName, tc := range testcases {
		t.Run(caseName, func(t *testing.T) {
			actual := lengthOfLongestSubstring(tc.str)
			t.Logf("str: %v, expected: %v, actual: %v", tc.str, tc.expected, actual)

			if actual != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
