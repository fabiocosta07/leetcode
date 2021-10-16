package main

func main() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	return traverseTree(root)
}

func traverseTree(node *TreeNode) int {
	if node == nil {
		return 0
	}
	currentNode := node
	counter := 0
	leftCounter := 0
	rightCounter := 0
	counter++
	if currentNode.Left != nil {
		leftCounter = traverseTree(currentNode.Left)
	}
	if currentNode.Right != nil {
		rightCounter = traverseTree(currentNode.Right)
	}
	if leftCounter == 0 {
		counter += rightCounter
		return counter
	}
	if rightCounter == 0 {
		counter += leftCounter
		return counter
	}
	if leftCounter < rightCounter {
		counter += leftCounter
	} else {
		counter += rightCounter
	}
	return counter
}
