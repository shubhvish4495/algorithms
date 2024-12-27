package main

import "fmt"

// https://leetcode.com/problems/maximum-product-subarray/description/
func main() {
	nums := []int{2, -5, 2, -4, -3}
	fmt.Println(maxProduct(nums))
}

func maxProduct(nums []int) int {
	maxP := nums[0]
	minP := nums[0]
	result := maxP

	for i := 1; i < len(nums); i++ {
		curr := nums[i]
		tempMaxP := max(maxP*curr, minP*curr, curr)
		minP = min(minP*curr, maxP*curr, curr)

		maxP = tempMaxP
		result = max(maxP, result)
	}

	return result
}
