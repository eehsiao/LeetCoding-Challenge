// Write a program to find the n-th ugly number.
// Ugly numbers are positive numbers whose prime factors only include 2, 3, 5.
// Example:
// Input: n = 10
// Output: 12
// Explanation: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 is the sequence of the first 10 ugly numbers.
// Note:
// 1 is typically treated as an ugly number.
// n does not exceed 1690.
//    Hide Hint #1
// The naive approach is to call isUgly for every number until you reach the nth one. Most numbers are not ugly. Try to focus your effort on generating only the ugly ones.
//    Hide Hint #2
// An ugly number must be multiplied by either 2, 3, or 5 from a smaller ugly number.
//    Hide Hint #3
// The key is how to maintain the order of the ugly numbers. Try a similar approach of merging from three sorted lists: L1, L2, and L3.
//    Hide Hint #4
// Assume you have Uk, the kth ugly number. Then Uk+1 must be Min(L1 * 2, L2 * 3, L3 * 5).

package main

func nthUglyNumber(n int) int {
	var ugly, i2, i3, i5 int
	nums := make([]int, n)
	nums[0] = 1

	for i := 1; i < n; i++ {
		ugly = min(min(nums[i2]*2, nums[i3]*3), nums[i5]*5)
		nums[i] = ugly

		if ugly == nums[i2]*2 {
			i2++
		}
		if ugly == nums[i3]*3 {
			i3++
		}
		if ugly == nums[i5]*5 {
			i5++
		}
	}

	return nums[n-1]
}
