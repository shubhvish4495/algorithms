package main

import "fmt"

func singleNonDuplicate(nums []int) int {
	l, h := 0, len(nums)-1

	for l < h {
		mid := l + (h-l)/2

		if mid%2 != 0 {
			mid -= 1
		}

		if nums[mid] == nums[mid+1] {
			l = mid + 2
		} else {
			h = mid
		}
	}
	return nums[l]
}

// https://leetcode.com/problems/single-element-in-a-sorted-array/
func main() {
	nums := []int{1, 1, 2, 3, 3, 4, 4, 8, 8}
	fmt.Println(singleNonDuplicate(nums))
}
