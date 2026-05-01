package greedy

import "testing"

func Test_jumpGame(t *testing.T) {

	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "",
			nums: []int{1, 1, 1, 3},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jumpGame(tt.nums); got != tt.want {
				t.Errorf("jumpGame(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
