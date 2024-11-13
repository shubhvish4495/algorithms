package main

import "fmt"

func nextPermutation(nums []int) bool {
	// Find the largest index k such that nums[k] < nums[k + 1]
	k := len(nums) - 2
	for k >= 0 && nums[k] >= nums[k+1] {
		k--
	}

	// If no such index exists, the permutation is the last permutation
	if k < 0 {
		reverse(nums, 0, len(nums)-1)
		return false
	}

	// Find the largest index l greater than k such that nums[k] < nums[l]
	l := len(nums) - 1
	for l >= 0 && nums[l] <= nums[k] {
		l--
	}

	// Swap the value of nums[k] with that of nums[l]
	nums[k], nums[l] = nums[l], nums[k]

	// Reverse the sequence from k+1 to end
	reverse(nums, k+1, len(nums)-1)
	return true
}

// Helper function to reverse a portion of the slice
func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func main() {
	nums := []int{1, 3, 2}
	nextPermutation(nums)

	fmt.Println(nums)
}
