// We write the integers of A and B (in the order they are given) on two separate horizontal lines.
// Now, we may draw connecting lines: a straight line connecting two numbers A[i] and B[j] such that:
// A[i] == B[j];
// The line we draw does not intersect any other connecting (non-horizontal) line.
// Note that a connecting lines cannot intersect even at the endpoints: each number can only belong to one connecting line.
// Return the maximum number of connecting lines we can draw in this way.
// Example 1:
// 1 4 2
// |  \
// 1 2 4
// Input: A = [1,4,2], B = [1,2,4]
// Output: 2
// Explanation: We can draw 2 uncrossed lines as in the diagram.
// We cannot draw 3 uncrossed lines, because the line from A[1]=4 to B[2]=4 will intersect the line from A[2]=2 to B[1]=2.
// Example 2:
// Input: A = [2,5,1,2,5], B = [10,5,2,1,5,2]
// Output: 3
// Example 3:
// Input: A = [1,3,7,1,7,5], B = [1,9,2,5,1]
// Output: 2
// Note:
// 1 <= A.length <= 500
// 1 <= B.length <= 500
// 1 <= A[i], B[i] <= 2000
//    Hide Hint #1
// Think dynamic programming. Given an oracle dp(i,j) that tells us how many lines A[i:], B[j:] [the sequence A[i], A[i+1], ... and B[j], B[j+1], ...] are uncrossed, can we write this as a recursion?

package main

//Result : Time Limited Exceeded
func checkDfs(A []int, B []int, i, j int) int {
	if i >= len(A) || j >= len(B) {
		return 0
	}
	if A[i] == B[j] {
		return 1 + checkDfs(A, B, i+1, j+1)
	}
	return max(checkDfs(A, B, i+1, j), checkDfs(A, B, i, j+1))

}

func maxUncrossedLines(A []int, B []int) int {
	//Result : Time Limited Exceeded
	// return checkDfs(A, B, 0, 0)

	var dp [][]int
	dp = append(dp, make([]int, len(B)+1))
	dp = append(dp, make([]int, len(B)+1))
	flag := 1
	for i := len(A) - 1; i >= 0; i-- {
		for j := len(B) - 1; j >= 0; j-- {
			dp[flag][j] = iif(A[i] == B[j], dp[flag^1][j+1]+1, max(dp[flag][j+1], dp[flag^1][j])).(int)
		}
		flag ^= 1
	}
	return dp[flag^1][0]
}
