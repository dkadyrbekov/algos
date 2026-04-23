package bin_search

import (
	"reflect"
	"testing"
)

func Test_findClosestElements(t *testing.T) {

	tests := []struct {
		name   string
		nums   []int
		k      int
		target int
		want   []int
	}{
		{
			name:   "empty",
			nums:   []int{-100, 0, 1, 3848, 6118, 6125, 6742, 7176, 7179, 8000},
			k:      4,
			target: 5383,
			want:   []int{3848, 6118, 6125, 6742},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findClosestElements(tt.nums, tt.k, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findClosestElements(%v, %v, %v) = %v, want %v", tt.nums, tt.k, tt.target, got, tt.want)
			}
		})
	}
}

func Test_findClosest(t *testing.T) {

	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "one",
			nums:   []int{1},
			target: 100,
			want:   0,
		},
		{
			name:   "empty",
			nums:   []int{-1, 0, 1, 2, 3, 4, 5},
			target: 2,
			want:   3,
		},
		{
			name:   "empty",
			nums:   []int{-1, 0, 1, 3, 4, 5},
			target: 2,
			want:   2,
		},
		{
			name:   "empty",
			nums:   []int{-100, 0, 1, 3848, 6118, 6125, 6742, 7176, 7179, 8000},
			target: 5383,
			want:   4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findClosest(tt.nums, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findClosestElements(%v, %v) = %v, want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
