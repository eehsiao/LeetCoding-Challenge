// You're given strings J representing the types of stones that are jewels, and S representing the stones you have.  Each character in S is a type of stone you have.  You want to know how many of the stones you have are also jewels.
// The letters in J are guaranteed distinct, and all characters in J and S are letters. Letters are case sensitive, so "a" is considered a different type of stone from "A".
// Example 1:
// Input: J = "aA", S = "aAAbbbb"
// Output: 3
// Example 2:
// Input: J = "z", S = "ZZ"
// Output: 0
// Note:
// S and J will consist of letters and have length at most 50.
// The characters in J are distinct.
//    Hide Hint #1
// For each stone, check if it is a jewel.

package main

import "testing"

func Test_numJewelsInStones(t *testing.T) {
	type args struct {
		J string
		S string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				J: "aA",
				S: "aAAbbbb",
			},
			want: 3,
		},
		{
			name: "case 2",
			args: args{
				J: "z",
				S: "ZZ",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numJewelsInStones(tt.args.J, tt.args.S); got != tt.want {
				t.Errorf("%v numJewelsInStones() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
