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

func countLevels(node *TreeNode, counter *int) int {
	*counter++
	if node.Left != nil {
		countLevels(node.Left, counter)
	}
	if node.Right != nil {
		countLevels(node.Right, counter)
	}
	return *counter
}

func maxDepth(root *TreeNode) int {

	var stack []*TreeNode = make([]*TreeNode, 0)

	n := root
loop1:
	for {
		for n.Left != nil {
			stack = append(stack, n)
			n = n.Left
		}
		if n.Right != nil {
			stack = append(stack, n)
			n = n.Right
			continue loop1
		}
		n = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for n.Right != nil {
			n = n.Right
		}

	}

}
