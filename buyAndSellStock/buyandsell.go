package buyandsellstock

func maxProfit(prices []int) int {
	var max int
	buy := 0
	sell := 1
	for sell < len(prices) {
		m := prices[sell] - prices[buy]
		if m > max {
			max = m
		}
		if prices[buy] > prices[sell] {
			buy = sell
		}
		sell++
	}
	return max
}
