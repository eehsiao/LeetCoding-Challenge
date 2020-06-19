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

import "testing"

func Test_longestDupSubstring(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				S: "banana",
			},
			want: "ana",
		},
		{
			name: "case 2",
			args: args{
				S: "abcd",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestDupSubstring(tt.args.S); got != tt.want {
				t.Errorf("longestDupSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
