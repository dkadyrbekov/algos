package cyclic_sort

import "testing"

func Test_findMissingNumber(t *testing.T) {
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
			name: "normal",
			nums: []int{1, 4, 5, 6, 8, 2, 0, 7},
			want: 3,
		},
		{
			name: "normal",
			nums: []int{1, 0, 2, 3, 4, 5, 6, 8, 9, 7, 11},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMissingNumber(tt.nums); got != tt.want {
				t.Errorf("findMissingNumber(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
