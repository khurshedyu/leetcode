package main

import "testing"

func TestMiddleNode(t *testing.T) {

	listNode1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	listNode2 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}
	listNode3 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}

	testcases := map[string]struct {
		input    *ListNode
		expected *ListNode
	}{
		"case1": {input: listNode1, expected: listNode1.Next},
		"case2": {input: listNode2, expected: listNode2.Next.Next},
		"case3": {input: listNode3, expected: listNode3.Next.Next},
	}

	for caseName, tc := range testcases {
		t.Run(caseName, func(t *testing.T) {
			actual := middleNode(tc.input)
			t.Logf("input: %v, expected: %v, actual: %v", tc.input, tc.expected, actual)

			if actual.Val != tc.expected.Val {
				t.Errorf("Expected %d, got %d", tc.expected.Val, actual.Val)
			}
		})
	}
}
