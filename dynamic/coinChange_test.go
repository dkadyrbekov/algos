package dynamic

import "testing"

func Test_coinChange(t *testing.T) {
	tests := []struct {
		name  string
		coins []int
		total int
		want  int
	}{
		{
			name:  "test",
			coins: []int{270, 192, 421, 269, 133, 6, 187, 474},
			total: 500,
			want:  10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChange(tt.coins, tt.total); got != tt.want {
				t.Errorf("coinChange(%v, %v) = %v, want %v", tt.coins, tt.total, got, tt.want)
			}
		})
	}
}
