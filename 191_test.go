package leetcode

import (
	"strconv"
	"testing"
)

func Test_hammingWeight(t *testing.T) {
	type args struct {
		num uint32
	}
	n1, _ := strconv.ParseUint("00000000000000000000000000001011", 2, 32)
	n2, _ := strconv.ParseUint("00000000000000000000000010000000", 2, 32)
	n3, _ := strconv.ParseUint("11111111111111111111111111111101", 2, 32)
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test", args: struct{ num uint32 }{num: uint32(n1)}, want: 3},
		{name: "test", args: struct{ num uint32 }{num: uint32(n2)}, want: 1},
		{name: "test", args: struct{ num uint32 }{num: uint32(n3)}, want: 31},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hammingWeight(tt.args.num); got != tt.want {
				t.Errorf("hammingWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
