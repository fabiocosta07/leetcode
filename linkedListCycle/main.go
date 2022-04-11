package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	var checkMap = make(map[*ListNode]bool)
	var currentNode *ListNode = head
	for currentNode.Next != nil {
		if _, ok := checkMap[currentNode.Next]; ok {
			return true
		}
		checkMap[currentNode.Next] = true
		currentNode = currentNode.Next
	}
	return false
}

func main() {

}
