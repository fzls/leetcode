package leetcode

import "testing"

func Test_isAnagram(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test", args: struct {
			s string
			t string
		}{s: "anagram", t: "nagaram"}, want: true},
		{name: "test", args: struct {
			s string
			t string
		}{s: "rat", t: "car"}, want: false},
		{name: "test", args: struct {
			s string
			t string
		}{s: "rat", t: "cart"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
