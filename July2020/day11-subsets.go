// Given a set of distinct integers, nums, return all possible subsets (the power set).
// Note: The solution set must not contain duplicate subsets.
// Example:
// Input: nums = [1,2,3]
// Output:
// [
//   [3],
//   [1],
//   [2],
//   [1,2,3],
//   [1,3],
//   [2,3],
//   [1,2],
//   []
// ]

package main

import (
	"sort"
)

func subsets(nums []int) [][]int {
	var results [][]int = [][]int{{}}

	for _, n := range nums {
		var sets [][]int = [][]int{}
		for _, s := range results {
			sets = append(sets, append([]int{n}, s...))
		}
		for _, s := range sets {
			results = append(results, s)
		}
	}
	for _, as := range results {
		sort.Ints(as)
	}

	return results
}
