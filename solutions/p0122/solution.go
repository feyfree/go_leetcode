package p0122

func maxProfit(prices []int) int {
	result := 0
	for i := 0; i < len(prices); i++ {
		if i+1 < len(prices) && prices[i+1] > prices[i] {
			result += prices[i+1] - prices[i]
		}
	}
	return result
}
