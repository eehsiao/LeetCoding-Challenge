// Given n, how many structurally unique BST's (binary search trees) that store values 1 ... n?
// Example:
// Input: 3
// Output: 5
// Explanation:
// Given n = 3, there are a total of 5 unique BST's:
//    1         3     3      2      1
//     \       /     /      / \      \
//      3     2     1      1   3      2
//     /     /       \                 \
//    2     1         2                 3

package main

func numTrees(n int) int {
	d := make([]int, n+1)
	d[0], d[1] = 1, 1

	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			d[i] += d[j-1] * d[i-j]
		}
	}

	return d[n]
}
