// Given an arbitrary ransom note string and another string containing letters from all the magazines, write a function that will return true if the ransom note can be constructed from the magazines ; otherwise, it will return false.
// Each letter in the magazine string can only be used once in your ransom note.
// Note:
// You may assume that both strings contain only lowercase letters.
// canConstruct("a", "b") -> false
// canConstruct("aa", "ab") -> false
// canConstruct("aa", "aab") -> true

package main

import (
	"strings"
)

func canConstruct(ransomNote string, magazine string) bool {
	mSlice := strings.Split(magazine, "")
	rSlice := strings.Split(ransomNote, "")
	m := make(map[string]int)

	for _, s := range mSlice {
		if _, ok := m[s]; ok {
			m[s]++
		} else {
			m[s] = 1
		}
	}
	for _, s := range rSlice {
		if _, ok := m[s]; !ok {
			return false
		}
		if m[s]--; m[s] < 0 {
			return false
		}
	}

	return true
}
