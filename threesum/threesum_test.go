package threesum

import (
	"slices"
	"testing"
)

func Test_threeSum(t *testing.T) {

	tests := []struct {
		name          string
		arg           []int
		wantSetsCount int
	}{
		{
			name:          "empy list",
			arg:           []int{},
			wantSetsCount: 0,
		},
		{
			name:          "one list",
			arg:           []int{1, 0, -1},
			wantSetsCount: 1,
		},
		{
			name:          "list of two",
			arg:           []int{1, -1},
			wantSetsCount: 0,
		},
		{
			name:          "no output",
			arg:           []int{1, 1, 9, 15, -3},
			wantSetsCount: 0,
		},
		{
			name:          "three output",
			arg:           []int{15, 9, 10, -1, 1, 0, -24, -9},
			wantSetsCount: 4,
		},
		{
			name:          "resolve unique",
			arg:           []int{15, 9, 10, -1, 1, 0, -24, -9, -9, -1, 1, 0},
			wantSetsCount: 4,
		},
		{
			name:          "regular case",
			arg:           []int{-1, 0, 1, 2, -1, -4},
			wantSetsCount: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := threeSum(tt.arg)

			if !uniqueSets(got) {
				t.Errorf("thressum() = %v, has not unique sets", got)
			}
			if len(got) != tt.wantSetsCount {
				t.Errorf("threeSum() = %v, wantedSetsCount %v", got, tt.wantSetsCount)
			}
		})
	}
}

func uniqueSets(sets [][]int) bool {
	for i := 0; i < len(sets)-1; i++ {
		found := findSet(sets[i], sets[i+1:])
		if found {
			return false
		}
	}

	return true
}

func findSet(searchSet []int, sets [][]int) (found bool) {
	slices.Sort(searchSet)
	for i := 0; i < len(sets); i++ {
		slices.Sort(sets[i])

		if len(searchSet) != len(sets[i]) {
			continue
		}

		for j := 0; j < len(searchSet); j++ {
			if searchSet[j] != sets[i][j] {
				break
			}

			if j == len(searchSet)-1 {
				return true
			}
		}
	}

	return false
}
