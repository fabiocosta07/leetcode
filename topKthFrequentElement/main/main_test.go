package main

import (
	"reflect"
	"testing"
)

func TestKth1(t *testing.T) {
	sample := []int{1, 1, 1, 2, 2, 3}
	k := 2

	res := topKFrequent(sample, k)

	if !reflect.DeepEqual(res, []int{1, 2}) {
		t.Errorf("Error")
	}
}

func TestKth2(t *testing.T) {
	sample := []int{1}
	k := 1

	res := topKFrequent(sample, k)

	if !reflect.DeepEqual(res, []int{1}) {
		t.Errorf("Error")
	}
}
