package main

import (
	"fmt"
	"slices"
)

func kthEle(a, b []int, k int) int {

	//base case
	if len(a) == 0 {
		return b[k]
	}

	if len(b) == 0 {
		return a[k]
	}

	for j := range b {
		l := j
		h := len(a) - 1
		e := findIndex(l, h, j, a, b)
		if e == -1 {
			a = slices.Insert(a, 0, b[j])
		} else {
			a = slices.Insert(a, e, b[j])
		}
	}

	return a[k-1]

}

func findIndex(l, h, j int, a, b []int) int {
	for l < h {
		if a[h] < b[j] {
			return h + 1
		}
		mid := l + (h-l)/2
		if a[mid] <= b[j] && a[mid+1] > b[j] {
			fmt.Println(b[j], mid)
			return mid
		}

		if a[mid] > b[j] {
			h = mid - 1
		} else {
			l = mid + 1
		}
	}

	return -1
}

// https://www.geeksforgeeks.org/problems/k-th-element-of-two-sorted-array1317/1
func main() {
	a := []int{2, 3, 6, 7, 9}
	b := []int{1, 4, 8, 10}
	k := 5
	fmt.Println(kthEle(a, b, k))
}
