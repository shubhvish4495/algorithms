package main

import (
	"fmt"
	"sort"
)

func jobSequencing(id, deadline, profit []int) []int {
	var maxProfit int
	sort.Slice(deadline, func(i, j int) bool {
		id[i], id[j] = id[j], id[i]
		profit[i], profit[j] = profit[j], profit[i]
		return deadline[i] < deadline[j]
	})

	maxProfit = profit[0]
	prevDeadline := deadline[0]
	currProfit := profit[0]
	count := 1
	for i := 1; i < len(deadline); i++ {
		if deadline[i] == prevDeadline {
			if currProfit < profit[i] {
				maxProfit -= currProfit
				currProfit = profit[i]
				maxProfit += currProfit
			}
		} else {
			maxProfit += profit[i]
			prevDeadline = deadline[i]
			count += 1
		}
	}

	return []int{count, maxProfit}
}

func main() {
	id := []int{1, 2, 3, 4, 5}
	deadline := []int{2, 1, 2, 1, 1}
	profit := []int{100, 19, 27, 25, 15}

	fmt.Println(jobSequencing(id, deadline, profit))
}
