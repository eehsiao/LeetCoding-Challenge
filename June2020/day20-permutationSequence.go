// The set [1,2,3,...,n] contains a total of n! unique permutations.
// By listing and labeling all of the permutations in order, we get the following sequence for n = 3:
// "123"
// "132"
// "213"
// "231"
// "312"
// "321"
// Given n and k, return the kth permutation sequence.
// Note:
// Given n will be between 1 and 9 inclusive.
// Given k will be between 1 and n! inclusive.
// Example 1:
// Input: n = 3, k = 3
// Output: "213"
// Example 2:
// Input: n = 4, k = 9
// Output: "2314"

package main

import (
	"fmt"
	"strconv"
)

func getPermutation(n int, k int) string {
	f, nums := make([]int, n), make([]int, 0)
	s := ""

	f[0] = 1
	nums = append(nums, 1)
	for i := 1; i < n; i++ {
		f[i] = f[i-1] * i
		nums = append(nums, i+1)
	}
	k--
	fmt.Println(f, nums)

	for i := n - 1; i >= 0; i-- {
		idx := k / f[i]
		k -= idx * f[i]

		s += strconv.Itoa(nums[idx])
		nums = append(nums[:idx], nums[idx+1:]...)
	}

	return s
}
