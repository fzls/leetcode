package leetcode

import "testing"

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test", args: struct{ num int }{num: 3}, want: "III"},
		{name: "test", args: struct{ num int }{num: 4}, want: "IV"},
		{name: "test", args: struct{ num int }{num: 9}, want: "IX"},
		{name: "test", args: struct{ num int }{num: 58}, want: "LVIII"},
		{name: "test", args: struct{ num int }{num: 1994}, want: "MCMXCIV"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.args.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
