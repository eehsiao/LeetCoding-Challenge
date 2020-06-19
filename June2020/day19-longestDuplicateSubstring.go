// Given a string S, consider all duplicated substrings: (contiguous) substrings of S that occur 2 or more times.  (The occurrences may overlap.)
// Return any duplicated substring that has the longest possible length.  (If S does not have a duplicated substring, the answer is "".)
// Example 1:
// Input: "banana"
// Output: "ana"
// Example 2:
// Input: "abcd"
// Output: ""
// Note:
// 2 <= S.length <= 10^5
// S consists of lowercase English letters.
//    Hide Hint #1
// Binary search for the length of the answer. (If there's an answer of length 10, then there are answers of length 9, 8, 7, ...)
//    Hide Hint #2
// To check whether an answer of length K exists, we can use Rabin-Karp 's algorithm.

package main

import "math"

func search(L, a, m, n int, nums []int) int {
	h, aL := 0, 1
	s := make(map[int]struct{}, 0)
	for i := 0; i < L; i++ {
		h = (h*a + nums[i]) % m
	}
	s[h] = struct{}{}
	for i := 1; i <= L; i++ {
		aL = (aL * a) % m
	}

	for start := 1; start < n-L+1; start++ {
		h = (h*a - nums[start-1]*aL%m + m) % m
		h = (h + nums[start+L-1]) % m
		if _, ok := s[h]; ok {
			return start
		}
		s[h] = struct{}{}
	}
	return -1
}

func longestDupSubstring(S string) string {
	sSlice := []byte(S)
	n, a, m := len(sSlice), 26, int(math.Pow(2, 32))
	l, r, L := 1, n, 0
	nums := make([]int, n)

	for i := 0; i < n; i++ {
		nums[i] = int(sSlice[i] - byte('a'))
	}

	for l <= r {
		L = l + (r-l)/2
		if search(L, a, m, n, nums) != -1 {
			l = L + 1
		} else {
			r = L - 1
		}
	}
	s := search(l-1, a, m, n, nums)
	return string(sSlice[s : s+l-1])
}
