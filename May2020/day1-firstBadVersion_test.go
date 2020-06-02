/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
package main

import "testing"

func Test_firstBadVersion(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				n: 5,
			},
			want: 4,
		},
		{
			name: "case 2",
			args: args{
				n: 5,
			},
			want: 3,
		},
		{
			name: "case 3",
			args: args{
				n: 5,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			badVer = tt.want
			if got := firstBadVersion(tt.args.n); got != tt.want {
				t.Errorf("%v firstBadVersion() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
