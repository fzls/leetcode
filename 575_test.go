package leetcode

import "testing"

func Test_distributeCandies(t *testing.T) {
	type args struct {
		candies []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test", args: struct{ candies []int }{candies: []int{1, 1, 2, 2, 3, 3}}, want: 3},
		{name: "test", args: struct{ candies []int }{candies: []int{1, 1, 2, 3}}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distributeCandies(tt.args.candies); got != tt.want {
				t.Errorf("distributeCandies(%v) = %v, want %v", tt.args.candies, got, tt.want)
			}
		})
	}
}
