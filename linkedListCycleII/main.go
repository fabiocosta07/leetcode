package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	var checkMap = make(map[*ListNode]bool)
	var currentNode *ListNode = head
	for currentNode != nil {
		if _, ok := checkMap[currentNode]; ok {
			return currentNode
		}
		checkMap[currentNode] = true
		currentNode = currentNode.Next
	}
	return nil
}

func main() {

}
