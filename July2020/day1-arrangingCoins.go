// You have a total of n coins that you want to form in a staircase shape, where every k-th row must have exactly k coins.
// Given n, find the total number of full staircase rows that can be formed.
// n is a non-negative integer and fits within the range of a 32-bit signed integer.
// Example 1:
// n = 5
// The coins can form the following rows:
// ¤
// ¤ ¤
// ¤ ¤
// Because the 3rd row is incomplete, we return 2.
// Example 2:
// n = 8
// The coins can form the following rows:
// ¤
// ¤ ¤
// ¤ ¤ ¤
// ¤ ¤
// Because the 4th row is incomplete, we return 3.

package main

import "math"

func arrangeCoins(n int) int {
	var l, r, k, c int = 0, n, 0, 0
	for l <= r {
		k = l + (r-l)/2
		c = k * (k + 1) / 2

		if c == n {
			return k
		}

		if n < c {
			r = k - 1
		} else {
			l = k + 1
		}
	}
	return r
}

func arrangeCoins_MathSqrt(n int) int {
	return int(math.Sqrt(2*float64(n)+0.25) - 0.5)
}
