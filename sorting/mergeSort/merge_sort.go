package main

import "fmt"

func mergeSort(nums []int) []int {

	if len(nums) < 2 {
		return nums
	}

	h := len(nums) - 1
	mid := (h) / 2

	lSorted := mergeSort(nums[:mid+1])
	rSorted := mergeSort(nums[mid+1:])
	return merge(lSorted, rSorted)
}

func merge(lSorted, rSorted []int) []int {
	if len(lSorted) == 0 {
		return rSorted
	}

	if len(rSorted) == 0 {
		return lSorted
	}

	lIndx := 0
	rIndx := 0
	sorted := make([]int, 0)

	for lIndx < len(lSorted) && rIndx < len(rSorted) {
		if lSorted[lIndx] < rSorted[rIndx] {
			sorted = append(sorted, lSorted[lIndx])
			lIndx++
		} else {
			sorted = append(sorted, rSorted[rIndx])
			rIndx++
		}
	}

	if lIndx != len(lSorted) {
		sorted = append(sorted, lSorted[lIndx:]...)
	}

	if rIndx != len(rSorted) {
		sorted = append(sorted, rSorted[rIndx:]...)
	}

	return sorted
}

func main() {

	nums := []int{4, 3, 2, 1}
	fmt.Println(mergeSort(nums))
}
