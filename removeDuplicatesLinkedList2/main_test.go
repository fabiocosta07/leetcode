package main

import (
	"fmt"
	"testing"
)

func TestDeleteNodes(t *testing.T) {
	vals := []int{1, 1, 2, 3, 3}
	head := addNodes(vals)
	result := deleteDuplicates(head)

	mapResult := make(map[int]bool)
	for result != nil {
		fmt.Println(result.Val)
		_, ok := mapResult[result.Val]
		if ok {
			t.Fail()
		}
		mapResult[result.Val] = true
		result = result.Next
	}
}
