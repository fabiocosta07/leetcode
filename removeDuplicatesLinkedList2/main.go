package main

import "fmt"

func main() {
	fmt.Println("test")
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addNodes(vals []int) *ListNode {
	var head *ListNode

	var node *ListNode
	for _, v := range vals {
		if head == nil {
			head = &ListNode{
				Val: v,
			}
			node = head
		} else {
			node.Next = &ListNode{
				Val: v,
			}
			node = node.Next
		}
	}
	return head
}

func deleteDuplicates(head *ListNode) *ListNode {
	node := head
	mapCheck := make(map[int]bool)
	var previousNode *ListNode
	for node != nil {
		if _, ok := mapCheck[node.Val]; ok {
			previousNode.Next = node.Next
			node = node.Next
			continue
		}
		mapCheck[node.Val] = true
		previousNode = node
		node = node.Next
	}
	return head
}
