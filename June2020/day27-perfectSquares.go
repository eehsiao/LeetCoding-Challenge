// Given a positive integer n, find the least number of perfect square numbers (for example, 1, 4, 9, 16, ...) which sum to n.
// Example 1:
// Input: n = 12
// Output: 3
// Explanation: 12 = 4 + 4 + 4.
// Example 2:
// Input: n = 13
// Output: 2
// Explanation: 13 = 4 + 9.

package main

func numSquares(n int) int {
	d := make([]int, n+1)

	for x := 1; x <= n; x++ {
		minVal, y, sq := x, 1, 1
		for sq <= x {
			minVal = min(minVal, 1+d[x-sq])
			y++
			sq = y * y
		}
		d[x] = minVal
	}
	return d[n]
}
