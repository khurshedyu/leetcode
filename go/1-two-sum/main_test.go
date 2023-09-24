package main

import "testing"

func TestTwoSum(t *testing.T) {
	testcases := map[string]struct {
		nums     []int
		target   int
		expected []int
	}{
		"testcase1": {[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		"testcase2": {[]int{3, 2, 4}, 6, []int{1, 2}},
		"testcase3": {[]int{3, 3}, 6, []int{0, 1}},
		"testcase4": {[]int{3, 2, 3}, 6, []int{0, 2}},
		"testcase5": {[]int{0, 2, 3, 0}, 0, []int{0, 3}},
	}

	for caseName, tc := range testcases {
		t.Run(caseName, func(t *testing.T) {
			actual := twoSum(tc.nums, tc.target)
			t.Logf("nums: %v, target: %v, expected: %v, actual: %v", tc.nums, tc.target, tc.expected, actual)

			for i := 0; i < len(actual); i++ {
				if actual[i] != tc.expected[i] {
					t.Errorf("Expected %v, got %v", tc.expected, actual)
				}
			}
		})
	}
}
