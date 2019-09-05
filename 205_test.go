package leetcode

import "testing"

func Test_isIsomorphic(t *testing.T) {
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
		}{s: "egg", t: "add"}, want: true},
		{name: "test", args: struct {
			s string
			t string
		}{s: "foo", t: "bar"}, want: false},
		{name: "test", args: struct {
			s string
			t string
		}{s: "paper", t: "title"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIsomorphic(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isIsomorphic() = %v, want %v", got, tt.want)
			}
		})
	}
}
