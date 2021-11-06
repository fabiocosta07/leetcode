package main

//* Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	return traverseTree(root, 0, targetSum)
}

func traverseTree(node *TreeNode, currentSum int, targetSum int) bool {
	if node == nil {
		return false
	}
	currentNode := node
	currentSum += node.Val

	if currentNode.Left == nil && currentNode.Right == nil {
		if currentSum == targetSum {
			return true
		}
	}
	resultLeft := false
	resultRight := false
	if currentNode.Left != nil {
		resultLeft = traverseTree(currentNode.Left, currentSum, targetSum)
	}
	if currentNode.Right != nil {
		resultRight = traverseTree(currentNode.Right, currentSum, targetSum)
	}

	if resultLeft || resultRight {
		return true
	}
	return false
}
