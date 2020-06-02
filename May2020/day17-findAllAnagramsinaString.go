// Given a string s and a non-empty string p, find all the start indices of p's anagrams in s.
// Strings consists of lowercase English letters only and the length of both strings s and p will not be larger than 20,100.
// The order of output does not matter.
// Example 1:
// Input:
// s: "cbaebabacd" p: "abc"
// Output:
// [0, 6]
// Explanation:
// The substring with start index = 0 is "cba", which is an anagram of "abc".
// The substring with start index = 6 is "bac", which is an anagram of "abc".
// Example 2:
// Input:
// s: "abab" p: "ab"
// Output:
// [0, 1, 2]
// Explanation:
// The substring with start index = 0 is "ab", which is an anagram of "ab".
// The substring with start index = 1 is "ba", which is an anagram of "ab".
// The substring with start index = 2 is "ab", which is an anagram of "ab".

package main

func findAnagrams(s string, p string) []int {
	latters := make([]int, 26)
	pSlice, sSlice := []byte(p), []byte(s)
	r, j := len(pSlice), 0
	var result []int
	for _, c := range pSlice {
		latters[c-byte('a')]++
	}

	for i := 0; i < len(sSlice); i++ {
		for ; j < len(sSlice) && j-i < len(pSlice); j++ {
			if latters[sSlice[j]-byte('a')] > 0 {
				r--
			}
			latters[sSlice[j]-byte('a')]--
		}
		if r == 0 && j-i == len(pSlice) {
			result = append(result, i)
		}
		if latters[sSlice[i]-byte('a')] >= 0 {
			r++
		}
		latters[sSlice[i]-byte('a')]++
	}
	return result
}
