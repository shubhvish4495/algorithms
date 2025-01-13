package main

import "fmt"

func search(nums []int, target int) int {
	l := 0
	h := len(nums) - 1

	if len(nums) == 1 {
		if target != nums[0] {
			return -1
		} else {
			return 0
		}
	}

	for l <= h {
		mid := l + (h-l)/2
		if nums[mid] == target {
			return mid
		}
		// check if left is sorted
		if nums[l] < nums[mid] {
			if nums[l] <= target && nums[mid] > target {
				h = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[l] < target && nums[mid] >= target {
				h = mid - 1
			} else {
				l = mid + 1
			}
		}

	}
	return -1
}

// https://leetcode.com/problems/search-in-rotated-sorted-array/
func main() {
	nums := []int{3, 5, 1}
	target := 1

	fmt.Println(search(nums, target))
}
