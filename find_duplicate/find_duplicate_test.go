package find_duplicate

import "testing"

func Test_findDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "empty slice",
			nums: []int{},
			want: -1,
		},
		{
			name: "slice of len 1",
			nums: []int{1},
			want: -1,
		},
		{
			name: "slice of len 2",
			nums: []int{1, 1},
			want: 1,
		},
		{
			name: "ordinar slice",
			nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 3},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicate(tt.nums); got != tt.want {
				t.Errorf("findDuplicate(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
