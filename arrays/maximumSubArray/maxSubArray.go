package main

import "fmt"

// https://leetcode.com/problems/maximum-subarray/description/

func maxSubArray(nums []int) int {

	maxSum := nums[0]
	currSum := nums[0]

	for i := 1; i < len(nums); i++ {
		currSum = max(currSum+nums[i], nums[i])

		if maxSum < currSum {
			maxSum = currSum
		}
	}

	return maxSum
}

func main() {
	arr := []int{-1, -2}
	fmt.Println(maxSubArray(arr))
}
