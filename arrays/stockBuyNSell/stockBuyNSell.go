package main

import "fmt"

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
func maxProfit(prices []int) int {

	maxProfit := 0
	smallestTillNow := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < smallestTillNow {
			smallestTillNow = prices[i]
		}
		maxProfit = max(maxProfit, prices[i]-smallestTillNow)
	}

	return maxProfit
}

func main() {
	prices := []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices))
}
