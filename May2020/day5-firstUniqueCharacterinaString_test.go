// Given a string, find the first non-repeating character in it and return it's index. If it doesn't exist, return -1.
// Examples:
// s = "leetcode"
// return 0.
// s = "loveleetcode",
// return 2.
// Note: You may assume the string contain only lowercase letters.
// its like https://github.com/eehsiao/30-Day-LeetCoding-Challenge/blob/master/week4-firstUniqueNumber.go

package main

import "testing"

func Test_firstUniqChar(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				s: "leetcode",
			},
			want: 0,
		},
		{
			name: "case 2",
			args: args{
				s: "loveleetcode",
			},
			want: 2,
		},
		{
			name: "case 3",
			args: args{
				s: "",
			},
			want: -1,
		},
		{
			name: "case 4",
			args: args{
				s: "aabb",
			},
			want: -1,
		},
		{
			name: "case 5",
			args: args{
				s: "dddccdbba",
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstUniqChar(tt.args.s); got != tt.want {
				t.Errorf("%v firstUniqChar() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
