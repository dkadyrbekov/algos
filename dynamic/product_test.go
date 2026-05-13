package dynamic

import (
	"math"
	"testing"
)

func Test_maxProduct(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "empty",
			nums: []int{},
			want: math.MinInt32,
		},
		{
			name: "one",
			nums: []int{2},
			want: 2,
		},
		{
			name: "Positive",
			nums: []int{1, 2, 3, 4},
			want: 24,
		},
		//{
		//	name: "With zeroes",
		//	nums: []int{1, 2, 0, 3, 4, 6, 0, 3, 7},
		//	want: 72,
		//},
		{
			name: "With negative",
			nums: []int{1, 2, -1, 3, 4, 6, -1, 3, 1, -10},
			want: 2160,
		},
		//{
		//	name: "With negative",
		//	nums: []int{0, -5, -1, 4, 10, -6, 6, -3, -9, 8, 0, 9, 8, -1, 6, 7, 8, 2, -5, -9},
		//	want: 2160,
		//},
		{
			name: "Only negative",
			nums: []int{-1, -2, -3, -4},
			want: -1,
		},
		//{
		//	name: "With negative",
		//	nums: []int{-1, -1, -5, 7, 1, -3, -5, 0, -6, 2, -2, 7, 1, 7, 8, -3, 1, -9, -6, -9, 4, -1, -3, 0, 10, 10, 6, -1, 0, -8, 6, -6, -5, -10, -6, 0, 6, 1, -1, 10, -5, -2, 5, -10, 0, 3, -2, 3, -9, 2, 5, -7, 9, 0, -8, 0, 1, -3, 0, 2, -2, 8, 6, -9, 9, -1, -6, -7, 1, -1, -6, 0, 10, 8, 1, 6, -6, -8, 1, 10, 0, -5, -1, 4, 10, -6, 6, -3, -9, 8, 0, 9, 8, -1, 6, 7, 8, 2, -5, -9},
		//	want: 2160,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.nums); got != tt.want {
				t.Errorf("maxProduct(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
