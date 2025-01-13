package main

import "fmt"

func maxSumIncreasing(nums []int, dp map[int]int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	if val, ok := dp[nums[0]]; ok {
		return val
	}

	maxSum := -999999
	for i := 0; i < len(nums); i++ {
		currEle := nums[i]
		tmaxSum := nums[i]
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > currEle {
				cSum := maxSumIncreasing(nums[j:], dp)
				tmaxSum = max(tmaxSum, currEle+cSum)
			}
		}
		dp[currEle] = tmaxSum
		maxSum = max(maxSum, tmaxSum)
	}

	return maxSum
}

func main() {
	arr := []int{4, 2, 4, 7}

	s := maxSumIncreasing(arr, make(map[int]int))
	fmt.Println(s)
}
