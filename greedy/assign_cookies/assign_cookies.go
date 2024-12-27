package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	var contentChildren int
	sort.Slice(g, func(i, j int) bool {
		return g[i] > g[j]
	})

	sort.Slice(s, func(i, j int) bool {
		return s[i] > s[j]
	})

	gIndx := 0
	sIndx := 0
	for gIndx < len(g) && sIndx < len(s) {
		// fmt.Println(gIndx, len(g), sIndx, len(s), gIndx < len(g) || sIndx < len(s))
		if g[gIndx] > s[sIndx] {
			gIndx++
			continue
		}
		sIndx++
		gIndx++
		contentChildren += 1
	}

	return contentChildren
}

func main() {
	// g = [1,2,3], s = [1,1]
	g := []int{1, 2}
	s := []int{1, 2, 3}

	fmt.Println(findContentChildren(g, s))
}
