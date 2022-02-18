package main

import (
	"fmt"
	"testing"
)

func TestDeleteNodes(t *testing.T) {
	vals := []int{1, 1, 1, 2, 3}
	head := addNodes(vals)
	result := deleteDuplicates(head)

	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}
