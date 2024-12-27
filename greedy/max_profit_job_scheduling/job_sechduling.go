package main

import (
	"fmt"
	"sort"
)

func inRange(st1, et1, st2, et2 int) bool {
	if st2 < et1 {
		return true
	}

	return false
}

type custStruct struct {
	startTime int
	endTime   int
	profit    int
}

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	var maxProfit int
	indices := make([]custStruct, len(endTime))

	for i := range len(endTime) {
		indices[i] = custStruct{
			startTime: startTime[i],
			endTime:   endTime[i],
			profit:    profit[i],
		}
	}

	sort.Slice(indices, func(i, j int) bool {
		return indices[i].endTime < indices[j].endTime
	})

	fmt.Println(indices)

	maxProfit = indices[0].profit
	currStTime := indices[0].startTime
	currEndTime := indices[0].endTime
	currProfit := maxProfit
	for i := 1; i < len(endTime); i++ {
		if inRange(currStTime, currEndTime, indices[i].startTime, indices[i].endTime) {
			if currProfit < indices[i].profit {
				maxProfit -= currProfit
			} else {
				continue
			}
		}
		currProfit = indices[i].profit
		maxProfit += currProfit
		currStTime = indices[i].startTime
		currEndTime = indices[i].endTime
	}

	return maxProfit
}

// startTime = [1,2,3,4,6], endTime = [3,5,10,6,9], profit = [20,20,100,70,60] -> 150
// startTime = [1,2,3,3], endTime = [3,4,5,6], profit = [50,10,40,70] -> 120

func main() {
	startTime := []int{1, 2, 3, 4, 6}
	endTime := []int{3, 5, 10, 6, 9}
	profit := []int{20, 20, 100, 70, 60}
	fmt.Println(jobScheduling(startTime, endTime, profit))
}
