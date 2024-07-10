package program

func MaxProfit(prices []int) int {
	minPrice := 1000
	maxProfit := 0

	//{7, 1, 5, 3, 6, 4}
	for _, price := range prices {
		if price < minPrice { // 5 // 1 // 5 < 1 // 3 < 1 // 6 < 1
			minPrice = price // 1
		} else if price-minPrice > maxProfit { // 5 - 1 > 0 // 3 - 1 > 4 // 6 - 1 > 4
			maxProfit = price - minPrice // 5
		}
	}
	return maxProfit // Returns the maximum profit that can be made from the given prices.
}
