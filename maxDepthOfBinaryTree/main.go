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

func maxDepth(root *TreeNode) int {
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
	if currentNode.Left != nil {
		leftCounter = traverseTree(currentNode.Left)
	}
	if currentNode.Right != nil {
		rightCounter = traverseTree(currentNode.Right)
	}
	if leftCounter > rightCounter {
		counter = +leftCounter
	} else {
		counter = +rightCounter
	}
	counter++
	return counter
}
