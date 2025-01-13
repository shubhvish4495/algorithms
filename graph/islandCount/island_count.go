package main

import (
	"fmt"
	"slices"
)

func numIslands(grid [][]byte) int {
	visited := make(map[string]struct{})

	islandCount := 0
	for i := range grid {
		for j := range grid[i] {
			indx := fmt.Sprintf("%d%d", i, j)
			if _, isVisited := visited[indx]; isVisited {
				continue
			}
			if string(grid[i][j]) == "1" {
				traverse(grid, visited)
				islandCount++
			}
		}
	}

	return islandCount
}

func traverse(grid [][]byte, i, j int, visited map[string]struct{}) {

	left := i - 1
	right := i + 1
	down := j + 1
	ip := j - 1

}

// https://leetcode.com/problems/number-of-islands/description/
func main() {
	x := []int{1, 5, 3}
	slices.Sort(x)

	y := slices.Insert(x, 4, 10)

	fmt.Println(y)

}
