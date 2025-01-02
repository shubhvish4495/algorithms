package main

import (
	"fmt"

	"github.com/shubhvish4495/algorithms/helper"
)

func dfs(graph [][]int) []int {
	dfsTraversal := make([]int, 0)
	isVisited := make(map[int]struct{})
	q := helper.GetNewStack()
	q.Push(0)
	for q.Length() != 0 {
		e := q.Pop()
		if _, ok := isVisited[e]; ok {
			continue
		}
		for i := len(graph[e]) - 1; i >= 0; i-- {
			q.Push(graph[e][i])
		}
		dfsTraversal = append(dfsTraversal, e)
		isVisited[e] = struct{}{}
	}

	return dfsTraversal
}

// https://www.geeksforgeeks.org/problems/depth-first-traversal-for-a-graph/1
func main() {
	graph := [][]int{{2, 3, 1}, {0}, {0, 4}, {0}, {2}}
	fmt.Println(dfs(graph))
}
