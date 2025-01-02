package main

import "fmt"

func coreLogicCoinChange(amount int, coins []int, dp map[string]int) int {
	if amount == 0 {
		return 0
	}

	if amount < 0 {
		return -1
	}

	key := fmt.Sprintf("%d", amount)
	if val, ok := dp[key]; ok {
		return val
	}

	maxVal := 999999999
	for i := range coins {
		tVal := coreLogicCoinChange(amount-coins[i], coins, dp)
		if tVal != -1 {
			maxVal = min(maxVal, tVal+1)
		}

	}

	if maxVal == 999999999 {
		maxVal = -1
	}

	dp[key] = maxVal
	return maxVal
}

func coinChange(coins []int, amount int) int {
	dp := make(map[string]int)
	// n := len(coins)
	return coreLogicCoinChange(amount, coins, dp)
}

func main() {

	coins := []int{1}
	amount := 0
	fmt.Println(coinChange(coins, amount))
}
