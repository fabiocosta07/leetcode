package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var previous *ListNode
	pointer := head
	for pointer != nil {
		next := pointer.Next
		pointer.Next = previous
		if next == nil {
			return pointer
		}
		previous = pointer
		pointer = next
	}
	return nil
}
