package bin_search

import "testing"

func Test_singleNonDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "empty",
			nums: []int{},
			want: -1,
		},
		{
			name: "one",
			nums: []int{1},
			want: 1,
		},
		{
			name: "three",
			nums: []int{1, 1, 2},
			want: 2,
		},
		{
			name: "even",
			nums: []int{1, 1, 2, 2, 5, 5, 6, 6, 7, 8, 8},
			want: 7,
		},
		{
			name: "odd",
			nums: []int{1, 1, 2, 2, 5, 5, 6, 6, 7, 8, 8, 9, 9},
			want: 7,
		},
		{
			name: "first",
			nums: []int{1, 2, 2, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9},
			want: 1,
		},
		{
			name: "last",
			nums: []int{1, 1, 2, 2, 5, 5, 6, 6, 7, 7, 8, 8, 9},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNonDuplicate(tt.nums); got != tt.want {
				t.Errorf("singleNonDuplicate(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
