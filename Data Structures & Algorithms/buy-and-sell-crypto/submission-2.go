func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		maxProfit = max(maxProfit, prices[i] - minPrice)
		if minPrice > prices[i] {
			minPrice = prices[i]
		}
	}
	return maxProfit
}
