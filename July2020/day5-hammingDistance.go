// The Hamming distance between two integers is the number of positions at which the corresponding bits are different.
// Given two integers x and y, calculate the Hamming distance.
// Note:
// 0 ≤ x, y < 231.
// Example:
// Input: x = 1, y = 4
// Output: 2
// Explanation:
// 1   (0 0 0 1)
// 4   (0 1 0 0)
//        ↑   ↑
// The above arrows point to positions where the corresponding bits are different.

package main

func hammingDistance(x int, y int) int {
	xor, dist := x^y, 0

	for xor != 0 {
		if xor%2 == 1 {
			dist++
		}
		xor = xor >> 1
	}

	return dist
}
