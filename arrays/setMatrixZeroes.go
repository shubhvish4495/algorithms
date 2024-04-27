package main

import "fmt"

// leetcode problem link -> https://leetcode.com/problems/set-matrix-zeroes/description/

func main() {
	a := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	// a := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}

	// zeroColRowVal :=
	zeroRowVal := make(map[int]struct{})
	zeroColVal := make(map[int]struct{})

	for i := range a {
		for j := range a[i] {
			if a[i][j] == 0 {
				zeroRowVal[i] = struct{}{}
				zeroColVal[j] = struct{}{}
			}
		}
	}

	for i := range zeroRowVal {
		for j := range a[i] {
			a[i][j] = 0
		}
	}

	for j := range zeroColVal {
		for i := range len(a) {
			a[i][j] = 0
		}
	}

	fmt.Print(a)
}
