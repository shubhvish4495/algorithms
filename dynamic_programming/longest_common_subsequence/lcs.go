package main

import "fmt"

func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make(map[string]int)
	return find(text1, text2, 0, 0, dp)
}

func find(s1, s2 string, i1, i2 int, dp map[string]int) int {
	if i1 == len(s1) || i2 == len(s2) {
		return 0
	}

	if val, ok := dp[fmt.Sprintf("%d%d", i1, i2)]; ok {
		return val
	}

	if s1[i1] == s2[i2] {
		dat := 1 + find(s1, s2, i1+1, i2+1, dp)
		dp[fmt.Sprintf("%d%d", i1, i2)] = dat
		return dat
	}

	dat := 0 + max(find(s1, s2, i1+1, i2, dp), find(s1, s2, i1, i2+1, dp))
	dp[fmt.Sprintf("%d%d", i1, i2)] = dat
	return dat
}

// https://leetcode.com/problems/longest-common-subsequence/description/
func main() {
	text1 := "abcde"
	text2 := "ace"

	fmt.Println(longestCommonSubsequence(text1, text2))
}
