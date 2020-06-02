// Given an arbitrary ransom note string and another string containing letters from all the magazines, write a function that will return true if the ransom note can be constructed from the magazines ; otherwise, it will return false.
// Each letter in the magazine string can only be used once in your ransom note.
// Note:
// You may assume that both strings contain only lowercase letters.
// canConstruct("a", "b") -> false
// canConstruct("aa", "ab") -> false
// canConstruct("aa", "aab") -> true

// bool canConstruct( string ransomNote, string magazine ) {
// 	unordered_map<char, int> magazineLetterFreq;
// 	for( auto c : magazine )  magazineLetterFreq[c]++;
// 	for( auto c : ransomNote )  if( --magazineLetterFreq[c] < 0 ) return false;
// 	return true;
// }

package main

import "testing"

func Test_canConstruct(t *testing.T) {
	type args struct {
		ransomNote string
		magazine   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				ransomNote: "a",
				magazine:   "b",
			},
			want: false,
		},
		{
			name: "case 2",
			args: args{
				ransomNote: "aa",
				magazine:   "ab",
			},
			want: false,
		},
		{
			name: "case 3",
			args: args{
				ransomNote: "aa",
				magazine:   "aab",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canConstruct(tt.args.ransomNote, tt.args.magazine); got != tt.want {
				t.Errorf("%v canConstruct() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
