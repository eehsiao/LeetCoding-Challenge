// Given a set of distinct positive integers, find the largest subset such that every pair (Si, Sj) of elements in this subset satisfies:
// Si % Sj = 0 or Sj % Si = 0.
// If there are multiple solutions, return any subset is fine.
// Example 1:
// Input: [1,2,3]
// Output: [1,2] (of course, [1,3] will also be ok)
// Example 2:
// Input: [1,2,4,8]
// Output: [1,2,4,8]

package main

import (
	"sort"
)

func largestDivisibleSubset(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	dp, maxSize, maxIdx := make([]int, len(nums)), -1, -1
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		subSize := 0
		for k := 0; k < i; k++ {
			if nums[i]%nums[k] == 0 && subSize < dp[k] {
				subSize = dp[k]
			}
		}

		dp[i] = subSize + 1

		if maxSize < dp[i] {
			maxSize = dp[i]
			maxIdx = i
		}
	}

	subSet := make([]int, 0)
	currSize, currTail := maxSize, nums[maxIdx]
	for i := maxIdx; i >= 0; i-- {
		if currSize == 0 {
			break
		}
		if currTail%nums[i] == 0 && currSize == dp[i] {
			// fmt.Println(nums[i], dp, maxSize, maxIdx)
			subSet = append([]int{nums[i]}, subSet...)
			currTail = nums[i]
			currSize--
		}
	}

	return subSet
}
