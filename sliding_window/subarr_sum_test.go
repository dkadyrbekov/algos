package sliding_window

import "testing"

func Test_minSubArrayLen(t *testing.T) {
	tests := []struct {
		name   string
		target int
		nums   []int
		want   int
	}{
		{
			name:   "empty",
			target: 0,
			nums:   []int{},
			want:   0,
		},
		{
			name:   "valid",
			target: 7,
			nums:   []int{2, 3, 1, 2, 4, 3},
			want:   2,
		},
		{
			name:   "valid",
			target: 10,
			nums:   []int{1, 2, 3, 4},
			want:   4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubArrayLen(tt.target, tt.nums); got != tt.want {
				t.Errorf("minSubArrayLen(%v, %v) = %v, want %v", tt.target, tt.nums, got, tt.want)
			}
		})
	}
}
