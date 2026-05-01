package greedy

import (
	"math"
	"sort"
)

func rescueBoats(people []int, limit int) int {
	sort.Ints(people)

	counter := 0
	i := 0
	for j := len(people) - 1; j >= i; j-- {
		pair, found := binSearch(people[i:j], limit-people[j])
		if found {
			people[pair], people[i] = people[i], people[pair]
			i++
		}
		counter++
	}

	return counter
}

func binSearch(people []int, limit int) (index int, found bool) {
	left := 0
	right := len(people) - 1

	maxPairWeight := math.MinInt

	for right >= left {
		middle := (right + left) / 2
		if people[middle] == limit {
			return middle, true
		}

		if people[middle] > limit {
			right = middle - 1
		} else {
			left = middle + 1

			if people[middle] > maxPairWeight {
				maxPairWeight = people[middle]
				index = middle
			}
		}
	}

	return index, maxPairWeight != math.MinInt
}
