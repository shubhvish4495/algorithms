package main

import (
	"fmt"

	"github.com/shubhvish4495/algorithms/helper"
)

func findKthLargest(nums []int, k int) int {
	h := helper.GetNewMinHeap(k)
	for i := 0; i < k && i < len(nums); i++ {
		h.Insert(nums[i])
	}

	// Process remaining elements
	for i := k; i < len(nums); i++ {
		if nums[i] > h.Peek() {
			h.Delete(0)
			h.Insert(nums[i])
		}
	}
	return h.Peek()
}

// https://leetcode.com/problems/kth-largest-element-in-an-array/description/
func main() {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k := 4

	fmt.Println(findKthLargest(nums, k))
}
