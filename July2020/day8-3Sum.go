// Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.
// Note:
// The solution set must not contain duplicate triplets.
// Example:
// Given array nums = [-1, 0, 1, 2, -1, -4],
// A solution set is:
// [
//   [-1, 0, 1],
//   [-1, -1, 2]
// ]
//    Hide Hint #1
// So, we essentially need to find three numbers x, y, and z such that they add up to the given value. If we fix one of the numbers say x, we are left with the two-sum problem at hand!
//    Hide Hint #2
// For the two-sum problem, if we fix one of the numbers, say
// x
// , we have to scan the entire array to find the next number
// y
// which is
// value - x
// where value is the input parameter. Can we change our array somehow so that this search becomes faster?
//    Hide Hint #3
// The second train of thought for two-sum is, without changing the array, can we use additional space somehow? Like maybe a hash map to speed up the search?

package main

import (
	"sort"
	"strconv"
)

func threeSum(nums []int) [][]int {
	var (
		ans    [][]int
		ansMap map[string]struct{} = make(map[string]struct{})
		n      int                 = len(nums)
	)

	sort.Ints(nums)

	for i := 0; i < n-2; i++ {
		val, l, r := nums[i], i+1, n-1
		for l < r {
			csum := val + nums[l] + nums[r]
			if csum == 0 {
				s := strconv.Itoa(val) + strconv.Itoa(nums[l]) + strconv.Itoa(nums[r])
				if _, ok := ansMap[s]; !ok {
					ansMap[s] = struct{}{}
					ans = append(ans, []int{val, nums[l], nums[r]})
				}
				l++
				r--
			} else if csum < 0 {
				l++
			} else {
				r--
			}
		}
	}

	return ans
}
