package main

import "fmt"

func sortColors(nums []int) {
	for {
		shouldContinue := false
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
				shouldContinue = true
			}
		}
		if !shouldContinue {
			break
		}
	}
}

func main() {
	arr := []int{2, 0, 2, 1, 1, 0}
	sortColors(arr)
	fmt.Println(arr)
}
