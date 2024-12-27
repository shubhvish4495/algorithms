package main

import (
	"fmt"
)

// https://leetcode.com/problems/longest-increasing-subsequence/
func lengthOfLIS(nums []int) int {
	lisArr := make([]int, len(nums))
	for i := range len(nums) {
		lisArr[i] = 1
	}

	maxLen := lisArr[0]

	for i := range len(nums) {
		iNum := nums[i]
		lisINum := lisArr[i]
		for j := i + 1; j < len(nums); j++ {
			jNum := nums[j]
			if jNum > iNum {
				lisArr[j] = max(lisArr[j], lisINum+1)
				maxLen = max(lisArr[j], maxLen)
			}
		}
	}
	return maxLen
}

func main() {
	nums := []int{0, 1, 0, 3, 2, 3}
	fmt.Println(lengthOfLIS(nums))
}
