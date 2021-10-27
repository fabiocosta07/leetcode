package buyandsellstock

import "testing"

func TestMaxProfit1(t *testing.T) {
	sample := []int{7, 1, 5, 3, 6, 4}
	res := maxProfit(sample)
	if res != 7 {
		t.Error("wrong result")
	}
}

func TestMaxProfit2(t *testing.T) {
	sample := []int{1, 2, 3, 4, 5}
	res := maxProfit(sample)
	if res != 4 {
		t.Error("wrong result")
	}
}

func TestMaxProfit3(t *testing.T) {
	sample := []int{7, 6, 4, 3, 1}
	res := maxProfit(sample)
	if res != 0 {
		t.Error("wrong result")
	}
}
