package main

import "fmt"

func binarySearch(arr []int, n int) int {
	l := 0
	h := len(arr) - 1

	for l <= h {
		mid := (l + h) / 2
		if arr[mid] == n {
			return mid
		}

		if arr[mid] > n {
			h = mid - 1
		} else {
			l = mid + 1
		}
	}

	return -1
}

func main() {
	arr := []int{1, 2, 4, 5, 6, 7, 8, 9}
	fmt.Println(binarySearch(arr, 90))
}
