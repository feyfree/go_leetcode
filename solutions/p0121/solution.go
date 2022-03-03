package p0121

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minVal := prices[0]
	profit := 0
	for i := 1; i < len(prices); i++ {
		if profit < prices[i]-minVal {
			profit = prices[i] - minVal
		}
		if prices[i] < minVal {
			minVal = prices[i]
		}
	}
	return profit
}
