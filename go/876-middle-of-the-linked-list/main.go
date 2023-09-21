package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	fmt.Println(middleNode(head))
}

func middleNode(head *ListNode) *ListNode {
	var slow = head
	var fast = head

	for {
		if fast == nil || fast.Next == nil {
			return slow
		}

		slow = slow.Next
		fast = fast.Next.Next
	}
}
