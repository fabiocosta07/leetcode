package buyandsellstock

import "testing"

func TestMaxProfit1(t *testing.T) {
	sample := []int{7, 1, 5, 3, 6, 4}
	res := maxProfit(sample)
	if res != 5 {
		t.Error("wrong result")
	}
}

func TestMaxProfit2(t *testing.T) {
	sample := []int{7, 6, 4, 3, 1}
	res := maxProfit(sample)
	if res != 0 {
		t.Error("wrong result")
	}
}
