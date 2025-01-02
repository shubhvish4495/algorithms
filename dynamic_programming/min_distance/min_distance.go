package main

import "fmt"

// https://leetcode.com/problems/edit-distance/description/

func minDistanceCore(word1, word2 string, dp map[string]int) int {
	if word1 == word2 {
		return 0
	}

	if word1 == "" && word2 != "" {
		return len(word2)
	}

	if word1 != "" && word2 == "" {
		return len(word1)
	}

	key := word1 + "_" + word2
	if val, ok := dp[key]; ok {
		return val
	}

	//for each word we can have 3 cases add remove or replace
	minVal := 99999999
	if word1[0] != word2[0] {
		removeVal := minDistanceCore(word1[1:], word2, dp) + 1
		replaceVal := minDistanceCore(word1[1:], word2[1:], dp) + 1
		addVal := minDistanceCore(word1, word2[1:], dp) + 1
		// fmt.Println(word1, word2, removeVal, replaceVal, addVal)
		minVal = min(minVal, removeVal, replaceVal, addVal)
	} else {
		minVal = min(minVal, minDistanceCore(word1[1:], word2[1:], dp))
	}
	dp[key] = minVal

	return minVal
}

func minDistance(word1 string, word2 string) int {
	dp := make(map[string]int)
	return minDistanceCore(word1, word2, dp)
}

func main() {
	w1 := "park"
	w2 := "spake"

	fmt.Println(minDistance(w1, w2))
}
