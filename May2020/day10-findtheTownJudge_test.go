// In a town, there are N people labelled from 1 to N.  There is a rumor that one of these people is secretly the town judge.
// If the town judge exists, then:
// The town judge trusts nobody.
// Everybody (except for the town judge) trusts the town judge.
// There is exactly one person that satisfies properties 1 and 2.
// You are given trust, an array of pairs trust[i] = [a, b] representing that the person labelled a trusts the person labelled b.
// If the town judge exists and can be identified, return the label of the town judge.  Otherwise, return -1.
// Example 1:
// Input: N = 2, trust = [[1,2]]
// Output: 2
// Example 2:
// Input: N = 3, trust = [[1,3],[2,3]]
// Output: 3
// Example 3:
// Input: N = 3, trust = [[1,3],[2,3],[3,1]]
// Output: -1
// Example 4:
// Input: N = 3, trust = [[1,2],[2,3]]
// Output: -1
// Example 5:
// Input: N = 4, trust = [[1,3],[1,4],[2,3],[2,4],[4,3]]
// Output: 3
// Note:
// 1 <= N <= 1000
// trust.length <= 10000
// trust[i] are all different
// trust[i][0] != trust[i][1]
// 1 <= trust[i][0], trust[i][1] <= N

package main

import "testing"

func Test_findJudge(t *testing.T) {
	type args struct {
		N     int
		trust [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				N:     2,
				trust: [][]int{{1, 2}},
			},
			want: 2,
		},
		{
			name: "case 2",
			args: args{
				N:     3,
				trust: [][]int{{1, 3}, {2, 3}},
			},
			want: 3,
		},
		{
			name: "case 3",
			args: args{
				N:     3,
				trust: [][]int{{1, 3}, {2, 3}, {3, 1}},
			},
			want: -1,
		},
		{
			name: "case 4",
			args: args{
				N:     3,
				trust: [][]int{{1, 2}, {2, 3}},
			},
			want: -1,
		},
		{
			name: "case 5",
			args: args{
				N:     4,
				trust: [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}, {4, 3}},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findJudge(tt.args.N, tt.args.trust); got != tt.want {
				t.Errorf("findJudge() = %v, want %v", got, tt.want)
			}
		})
	}
}
