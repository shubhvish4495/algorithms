package main

import (
	"fmt"

	"github.com/shubhvish4495/algorithms/helper"
)

func bfs(matrix [][]int) []int {
	var bfsTraversal []int
	q := helper.GetNewQueue()
	visited := make(map[int]struct{})

	q.Push(0)

	for q.Length() != 0 {
		ele := q.Pop()
		// already visited skipping
		if _, ok := visited[ele]; ok {
			continue
		}

		// fmt.Println(ele)
		bfsTraversal = append(bfsTraversal, ele)
		for i := range matrix[ele] {
			q.Push(matrix[ele][i])
		}
		visited[ele] = struct{}{}
	}

	return bfsTraversal
}

func main() {
	matrix := [][]int{{1, 2}, {0, 2}, {0, 1, 3, 4}, {2}, {2}}
	fmt.Println(bfs(matrix))
}
