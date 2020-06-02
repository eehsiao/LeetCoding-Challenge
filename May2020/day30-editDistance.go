// Given two words word1 and word2, find the minimum number of operations required to convert word1 to word2.
// You have the following 3 operations permitted on a word:
// Insert a character
// Delete a character
// Replace a character
// Example 1:
// Input: word1 = "horse", word2 = "ros"
// Output: 3
// Explanation:
// horse -> rorse (replace 'h' with 'r')
// rorse -> rose (remove 'r')
// rose -> ros (remove 'e')
// Example 2:
// Input: word1 = "intention", word2 = "execution"
// Output: 5
// Explanation:
// intention -> inention (remove 't')
// inention -> enention (replace 'i' with 'e')
// enention -> exention (replace 'n' with 'x')
// exention -> exection (replace 'n' with 'c')
// exection -> execution (insert 'u')

package main

import "strings"

func minDistance(word1 string, word2 string) int {
	n, m := len(word1), len(word2)
	w1Split, w2Split := strings.Split(word1, ""), strings.Split(word2, "")
	dp := make([][]int, n+1)
	if n == 0 || m == 0 {
		return n + m
	}

	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
		dp[i][0] = i
	}
	for i := 0; i <= m; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			l, d := dp[i-1][j]+1, dp[i][j-1]+1
			ld := dp[i-1][j-1]
			if w1Split[i-1] != w2Split[j-1] {
				ld++
			}
			dp[i][j] = min(l, min(d, ld))
		}
	}

	return dp[n][m]
}
