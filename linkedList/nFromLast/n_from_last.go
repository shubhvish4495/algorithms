package main

import (
	"fmt"

	ll "github.com/shubhvish4495/algorithms/linkedList"
)

func removeNthFromEnd(head *ll.ListNode, n int) *ll.ListNode {
	temp := head
	nodeCount := 0
	for temp != nil {
		nodeCount += 1
		temp = temp.Next
	}

	// nth node from end will be nodeCount - n
	ncFromBeginning := nodeCount - n
	temp = head

	if ncFromBeginning == 0 {
		head = head.Next
		return head
	}

	for ncFromBeginning != 0 {
		ncFromBeginning -= 1
		if ncFromBeginning == 0 {
			temp.Next = temp.Next.Next
		}
		temp = temp.Next
	}
	return head
}

func traverse(head *ll.ListNode) {
	temp := head
	for temp != nil {
		fmt.Print(temp.Val)
		temp = temp.Next
	}

	fmt.Println()
}

func main() {

	ll := &ll.ListNode{
		Val: 1,
		Next: &ll.ListNode{
			Val: 2,
		},
	}

	n := 2
	traverse(ll)
	removeNthFromEnd(ll, n)
	traverse(ll)

}
