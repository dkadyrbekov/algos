package circular_array_loop

import "testing"

func Test_circularArrayLoop(t *testing.T) {

	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "empty slice",
			nums: []int{},
			want: false,
		},
		{
			name: "slice length 1",
			nums: []int{-4852},
			want: false,
		},
		{
			name: "cycle of 1 element",
			nums: []int{3, 1, 1},
			want: false,
		},
		{
			name: "cycle of 3 elements",
			nums: []int{1, 1, 1},
			want: true,
		},
		{
			name: "cycle of 3 elements",
			nums: []int{2, 1, 2},
			want: true,
		},
		{
			name: "cycle of negative 3 elements",
			nums: []int{-2, -1, 2},
			want: true,
		},
		{
			name: "no cycle",
			nums: []int{-2, -2, 2},
			want: false,
		},
		{
			name: "tricky cycle",
			nums: []int{3, 3, 1, -1, 2},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := circularArrayLoop(tt.nums); got != tt.want {
				t.Errorf("circularArrayLoop(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
