// Given two strings s1 and s2, write a function to return true if s2 contains the permutation of s1. In other words, one of the first string's permutations is the substring of the second string.
// Example 1:
// Input: s1 = "ab" s2 = "eidbaooo"
// Output: True
// Explanation: s2 contains one permutation of s1 ("ba").
// Example 2:
// Input:s1= "ab" s2 = "eidboaoo"
// Output: False
// Note:
// The input strings only contain lower case letters.
// The length of both given strings is in range [1, 10,000].
//    Hide Hint #1
// Obviously, brute force will result in TLE. Think of something else.
//    Hide Hint #2
// How will you check whether one string is a permutation of another string?
//    Hide Hint #3
// One way is to sort the string and then compare. But, Is there a better way?
//    Hide Hint #4
// If one string is a permutation of another string then they must one common metric. What is that?
//    Hide Hint #5
// Both strings must have same character frequencies, if one is permutation of another. Which data structure should be used to store frequencies?
//    Hide Hint #6
// What about hash table? An array of size 26?

package main

import (
	"reflect"
	"sort"
	"strings"
)

func checkInclusion(s1 string, s2 string) bool {
	s1Slice, s2Slice := []byte(s1), []byte(s2)
	if len(s1Slice) <= len(s2Slice) {
		s1Map, s2Map, cnt := make([]byte, 26), make([]byte, 26), 0
		for i := 0; i < len(s1Slice); i++ {
			s1Map[s1Slice[i]-byte('a')]++
			s2Map[s2Slice[i]-byte('a')]++
		}
		for i := 0; i < 26; i++ {
			if s1Map[i] == s2Map[i] {
				cnt++
			}
		}
		for i := 0; i < len(s2Slice)-len(s1Slice); i++ {
			r, l := s2Slice[i+len(s1Slice)]-byte('a'), s2Slice[i]-byte('a')
			if cnt == 26 {
				return true
			}
			s2Map[r]++
			if s2Map[r] == s1Map[r] {
				cnt++
			} else if s2Map[r] == s1Map[r]+1 {
				cnt--
			}
			s2Map[l]--
			if s2Map[l] == s1Map[l] {
				cnt++
			} else if s2Map[l] == s1Map[l]-1 {
				cnt--
			}
		}
		return cnt == 26
	}

	return false
}

// Time Limit Exceeded
func checkInclusion_timeout(s1 string, s2 string) bool {
	s1Slice := strings.Split(s1, "")
	s2Slice := strings.Split(s2, "")
	subS2 := make([]string, len(s1Slice))
	sort.Strings(s1Slice)
	for i := 0; i <= len(s2)-len(s1); i++ {
		copy(subS2, s2Slice[i:i+len(s1Slice)])
		if sort.Strings(subS2); reflect.DeepEqual(s1Slice, subS2) {
			return true
		}
	}
	return false
}
