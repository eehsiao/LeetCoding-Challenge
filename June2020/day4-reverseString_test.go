// Write a function that reverses a string. The input string is given as an array of characters char[].
// Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.
// You may assume all the characters consist of printable ascii characters.
// Example 1:
// Input: ["hello"]
// Output: ["olleh"]
// Example 2:
// Input: ["Hannah"]
// Output: ["hannaH"]
//    Hide Hint #1
// The entire logic for reversing a string is based on using the opposite directional two-pointer approach!

package main

import (
	"reflect"
	"testing"
)

func Test_reverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "case 1",
			args: args{
				s: []byte("hello"),
			},
			want: []byte("olleh"),
		},
		{
			name: "case 2",
			args: args{
				s: []byte("Hannah"),
			},
			want: []byte("hannaH"),
		},
		{
			name: "case 3",
			args: args{
				s: []byte("A man, a plan, a canal: Panama"),
			},
			want: []byte("amanaP :lanac a ,nalp a ,nam A"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if reverseString(tt.args.s); !reflect.DeepEqual(tt.args.s, tt.want) {
				t.Errorf("%v twoCitySchedCost() = %v, want %v", tt.name, string(tt.args.s), string(tt.want))
			}
		})
	}
}
