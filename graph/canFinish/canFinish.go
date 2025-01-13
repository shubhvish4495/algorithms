package main

import (
	"fmt"
)

func canFinish(numCourses int, prerequisites [][]int) bool {

	dataArr := [][]int{}
	for range numCourses {
		dataArr = append(dataArr, make([]int, numCourses))
	}

	for k := range prerequisites {
		i := prerequisites[k][0]
		j := prerequisites[k][1]

		dataArr[i][j] = 1
	}

	for i := range dataArr {
		for j := range dataArr[i] {
			visited := make(map[int]struct{})
			if dataArr[i][j] == 1 {
				if !canTraverse(j, dataArr, visited) {
					return false
				}
			}
		}

	}

	return true
}

func canTraverse(i int, arr [][]int, visited map[int]struct{}) bool {
	// visited[i] = struct{}{}
	for x := range arr[i] {
		if arr[i][x] == 1 {
			if _, ok := visited[x]; ok {
				return false
			}
			visited[x] = struct{}{}
			if !canTraverse(x, arr, visited) {
				return false
			}
		}
	}

	return true
}

// https://leetcode.com/problems/course-schedule/description/
func main() {
	prerequisites := [][]int{{0, 3}, {1, 3}, {2, 0}, {2, 1}, {3, 2}}
	numCourses := 5

	fmt.Println(canFinish(numCourses, prerequisites))
}
