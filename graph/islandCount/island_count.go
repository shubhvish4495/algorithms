package main

import (
	"fmt"
	"slices"
)

func numIslands(grid [][]byte) int {
	return 0
}

// https://leetcode.com/problems/number-of-islands/description/
func main() {
	x := []int{1, 5, 3}
	slices.Sort(x)

	y := slices.Insert(x, 4, 10)

	fmt.Println(y)

}
