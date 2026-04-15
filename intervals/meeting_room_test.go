package intervals

import "testing"

func Test_findSets(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      int
	}{
		{
			name:      "empty intervals",
			intervals: [][]int{},
			want:      0,
		},
		{
			name: "valid intervals",
			intervals: [][]int{
				{2, 8}, {3, 9}, {5, 11}, {8, 20}, {11, 17},
			},
			want: 3,
		},
		{
			name: "valid intervals",
			intervals: [][]int{
				{1, 2}, {1, 2}, {1, 2},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSets(tt.intervals); got != tt.want {
				t.Errorf("findSets(%v) = %v, want %v", tt.intervals, got, tt.want)
			}
		})
	}
}
