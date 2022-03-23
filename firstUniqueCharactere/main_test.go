package main

import (
	"testing"
)

func TestUnique1(t *testing.T) {
	input := "leetcode"

	r := firstUniqChar(input)
	if r != 0 {
		t.Errorf("Error")
	}
}

func TestUnique2(t *testing.T) {
	input := "loveleetcode"

	r := firstUniqChar(input)
	if r != 2 {
		t.Errorf("Error")
	}
}

func TestUnique3(t *testing.T) {
	input := "aabb"

	r := firstUniqChar(input)
	if r != -1 {
		t.Errorf("Error")
	}
}
