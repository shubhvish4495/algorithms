package main

import "fmt"

func coreKnapsack(remWeight, i int, w, p []int, dp map[string]int) int {
	if remWeight <= 0 || i >= len(w) {
		return 0
	}
	currW := w[i]
	currP := p[i]

	if val, ok := dp[fmt.Sprintf("%d%d", currW, currP)]; ok {
		return val
	}

	var wcVal, ncVal int
	//while considering if and only if currW <= remWeight
	if currW <= remWeight {
		wcVal = coreKnapsack(remWeight-w[i], i+1, w, p, dp) + p[i]
	}

	//while not considering
	ncVal = coreKnapsack(remWeight, i+1, w, p, dp)
	tempVal := max(wcVal, ncVal)
	dp[fmt.Sprintf("%d%d", currW, currP)] = tempVal
	return tempVal
}

func knapsackSol(w int, profit, weight []int) int {
	dp := make(map[string]int)
	return coreKnapsack(w, 0, weight, profit, dp)
}

// N = 3, W = 4, profit[] = {1, 2, 3}, weight[] = {4, 5, 1}
func main() {
	w := 4
	profit := []int{1, 2, 3}
	weight := []int{4, 5, 1}
	fmt.Println(knapsackSol(w, profit, weight))
}
