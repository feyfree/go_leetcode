package p0123

func maxProfit(prices []int) int {
	n := len(prices)
	buy1 := -prices[0]
	sell1 := 0
	buy2 := -prices[0]
	sell2 := 0
	for i := 1; i < n; i++ {
		if buy1 < -prices[i] {
			buy1 = -prices[i]
		}
		if sell1 < buy1+prices[i] {
			sell1 = buy1 + prices[i]
		}
		if buy2 < sell1-prices[i] {
			buy2 = sell1 - prices[i]
		}
		if sell2 < buy2+prices[i] {
			sell2 = buy2 + prices[i]
		}
	}
	return sell2
}
