package main

import (
	"fmt"
	"sort"
)

func MaxMeetPossible(start, end []int) int {

	if len(start) != len(end) {
		return 0
	}

	sort.Slice(end, func(i, j int) bool {
		start[i], start[j] = start[j], start[i]
		return end[i] < end[j]
	})

	fmt.Println(start, end)

	lastEndIndx := end[0]
	meetCount := 1
	for i := 1; i < len(start); i++ {
		if start[i] > lastEndIndx {
			meetCount++
			lastEndIndx = end[i]
		}
	}

	return meetCount
}

func main() {
	start := []int{25, 0, 14, 24, 18, 3, 17}
	end := []int{29, 25, 24, 26, 25, 23, 18}

	fmt.Println(MaxMeetPossible(start, end))

}
