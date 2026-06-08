package cyclic_sort

import "testing"

func Test_smallestMissingPositiveInteger(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "empty",
			nums: []int{},
			want: 1,
		},
		{
			name: "empty",
			nums: []int{-1, -4, 0, 4, 6, 7, 13, -1, 0, 0, 17},
			want: 1,
		},
		{
			name: "empty",
			nums: []int{55, 22, 52, 100, 1, 3, -5},
			want: 2,
		},
		{
			name: "empty",
			nums: []int{1, 2, 3, 4, 5, 6, 7},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestMissingPositiveInteger(tt.nums); got != tt.want {
				t.Errorf("smallestMissingPositiveInteger(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}

//func Test_smallestMissingPositiveInteger(t *testing.T) {
//	tests := []struct {
//		name string
//		nums []int
//		want int
//	}{
//		{
//			name: "empty",
//			nums: []int{},
//			want: -1,
//		},
//		{
//			name: "empty",
//			nums: []int{},
//			want: -1,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := smallestMissingPositiveInteger(tt.nums); got != tt.want {
//				t.Errorf("smallestMissingPositiveInteger(%v) = %v, want %v", tt.nums, got, tt.want)
//			}
//		})
//	}
//}
