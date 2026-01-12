package main

func maxProfit(prices []int) int {
	minPrice := prices[0]
	profit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if profit < prices[i]-minPrice {
			profit = prices[i] - minPrice
		}
	}

	return profit
}
