package main

import (
	"fmt"
	"math"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var log1 float64 = 0
	var sum1 float64
	for l1 != nil {
		sum1 += float64(l1.Val) * math.Pow(10, log1)
		l1 = l1.Next
		log1++
	}
	var log2 float64 = 0
	var sum2 float64
	for l2 != nil {
		sum2 += float64(l2.Val) * math.Pow(10, log2)
		l2 = l2.Next
		log2++
	}
	var sum = sum1 + sum2
	testSum := fmt.Sprint(int(sum))
	charArr := []rune(testSum)

	var nodeResult *ListNode
	var nodeCurrent *ListNode
	for i := len(charArr) - 1; i >= 0; i-- {
		val, _ := strconv.Atoi(string(charArr[i]))
		if nodeCurrent == nil {
			nodeCurrent = &ListNode{Val: val, Next: nil}
			nodeResult = nodeCurrent
			continue
		}
		nodeCurrent.Next = &ListNode{Val: val, Next: nil}
		nodeCurrent = nodeCurrent.Next
	}
	return nodeResult
}

func main() {

}
