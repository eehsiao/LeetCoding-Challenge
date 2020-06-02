// Given a string, sort it in decreasing order based on the frequency of characters.
// Example 1:
// Input:
// "tree"
// Output:
// "eert"
// Explanation:
// 'e' appears twice while 'r' and 't' both appear once.
// So 'e' must appear before both 'r' and 't'. Therefore "eetr" is also a valid answer.
// Example 2:
// Input:
// "cccaaa"
// Output:
// "cccaaa"
// Explanation:
// Both 'c' and 'a' appear three times, so "aaaccc" is also a valid answer.
// Note that "cacaca" is incorrect, as the same characters must be together.
// Example 3:
// Input:
// "Aabb"
// Output:
// "bbAa"
// Explanation:
// "bbaA" is also a valid answer, but "Aabb" is incorrect.
// Note that 'A' and 'a' are treated as two different characters.

package main

import "testing"

func Test_frequencySort(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		want1 string
		want2 string
	}{
		{
			name:  "case 1",
			args:  args{s: "tree"},
			want1: "eetr",
			want2: "eert",
		},
		{
			name:  "case 2",
			args:  args{s: "cccaaa"},
			want1: "cccaaa",
			want2: "aaaccc",
		},
		{
			name:  "case 3",
			args:  args{s: "Aabb"},
			want1: "bbAa",
			want2: "bbaA",
		},
		{
			name:  "case 4",
			args:  args{s: "2a554442f544asfasssffffasss"},
			want1: "sssssssffffff44444aaaa55522",
			want2: "sssssssffffff44444aaaa55522",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := frequencySort(tt.args.s); got != tt.want1 && got != tt.want2 {
				t.Errorf("%v frequencySort() = %v, want %v or %v", tt.name, got, tt.want1, tt.want2)
			}
		})
	}
}
