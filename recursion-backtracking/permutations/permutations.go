package main

import (
	"fmt"
	"slices"
)

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}
	iNum := nums[0]
	retPerms := permute(nums[1:])
	finalRetPerms := make([][]int, 0)
	for i := range retPerms {
		tempPerms := retPerms[i]
		for j := range len(tempPerms) + 1 {
			// t2 := make([]int, len(tempPerms))
			// copy(t2, tempPerms)
			finalRetPerms = append(finalRetPerms, slices.Insert(tempPerms, j, iNum))
		}
	}

	return finalRetPerms
}

// https://leetcode.com/problems/permutations/description/
func main() {
	nums := []int{5, 4, 6, 1}
	fmt.Println(permute(nums))
}
