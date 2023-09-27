package main

import "testing"

func TestRomanToInt(t *testing.T) {
	testcases := map[string]struct {
		s        string
		expected int
	}{
		"case1": {s: "III", expected: 3},
		"case2": {s: "IV", expected: 4},
		"case3": {s: "IX", expected: 9},
		"case4": {s: "LVIII", expected: 58},
		"case5": {s: "MCMXCIV", expected: 1994},
	}

	for caseName, tc := range testcases {
		t.Run(caseName, func(t *testing.T) {
			actual := romanToInt(tc.s)
			t.Logf("s: %v, expected: %v, actual: %v", tc.s, tc.expected, actual)

			if actual != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
