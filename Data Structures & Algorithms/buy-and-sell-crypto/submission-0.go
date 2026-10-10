func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func maxProfit(prices []int) int {
	minBuy := 0
	maxSale := 0
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		minBuy = min(minBuy, prices[i-1])
		maxProfit = max(maxProfit, prices[i] - minBuy)
	}
	for i := len(prices) - 2; i >= 0; i-- {
		maxSale = max(maxSale, prices[i+1])
		maxProfit = max(maxProfit, maxSale - prices[i])
	}
	return maxProfit
}
