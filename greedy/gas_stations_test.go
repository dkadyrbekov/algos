package greedy

import "testing"

func Test_gasStationJourney(t *testing.T) {
	tests := []struct {
		name string
		gas  []int
		cost []int
		want int
	}{
		{
			name: "empty",
			gas:  []int{},
			cost: []int{},
			want: -1,
		},
		{
			name: "one positive",
			gas:  []int{1},
			cost: []int{1},
			want: 0,
		},
		{
			name: "one negative",
			gas:  []int{2},
			cost: []int{3},
			want: -1,
		},
		{
			name: "several positive",
			gas:  []int{1, 2, 3, 4, 5},
			cost: []int{3, 4, 5, 1, 2},
			want: 3,
		},
		{
			name: "several positive",
			gas:  []int{1, 2, 3, 4, 5},
			cost: []int{3, 4, 5, 1, 2},
			want: 3,
		},
		{
			name: "several negative",
			gas:  []int{1, 2, 3, 4, 5},
			cost: []int{3, 100, 5, 1, 2},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gasStationJourney(tt.gas, tt.cost); got != tt.want {
				t.Errorf("gasStationJourney(%v, %v) = %v, want %v", tt.gas, tt.cost, got, tt.want)
			}
		})
	}
}
