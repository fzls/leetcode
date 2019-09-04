package leetcode

import (
	"strconv"
	"testing"
)

func Test_reverseBits(t *testing.T) {
	type args struct {
		num uint32
	}
	n1, _ := strconv.ParseUint("00000010100101000001111010011100", 2, 64)
	w1, _ := strconv.ParseUint("00111001011110000010100101000000", 2, 64)
	n2, _ := strconv.ParseUint("11111111111111111111111111111101", 2, 64)
	w2, _ := strconv.ParseUint("10111111111111111111111111111111", 2, 64)
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{name: "test", args: struct{ num uint32 }{num: uint32(n1)}, want: uint32(w1)},
		{name: "test", args: struct{ num uint32 }{num: uint32(n2)}, want: uint32(w2)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBits(tt.args.num); got != tt.want {
				t.Errorf("reverseBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
