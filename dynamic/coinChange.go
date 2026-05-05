package dynamic

import (
	"math"
)

func coinChange(coins []int, total int) int {
	var changes = make(map[int]int)

	return cChange(coins, total, changes)
}

func cChange(coins []int, total int, changes map[int]int) int {
	if total == 0 {
		changes[total] = 0
		return 0
	}

	minChange := math.MaxInt

	for _, c := range coins {
		if total-c < 0 {
			continue
		}

		subCh, ok := changes[total-c]
		if !ok {
			subCh = cChange(coins, total-c, changes)
		}

		if subCh == -1 {
			continue
		}

		if minChange > subCh+1 {
			minChange = subCh + 1
		}
	}

	if minChange == math.MaxInt {
		minChange = -1
	}

	changes[total] = minChange

	return minChange
}
