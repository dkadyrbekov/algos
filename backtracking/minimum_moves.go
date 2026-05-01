package backtracking

import "math"

type coordinate2 struct {
	i int
	j int
}

func minimumMoves(grid [][]int) int {
	zerosC := make([]coordinate2, 0)
	var extraC coordinate2

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				zerosC = append(zerosC, coordinate2{i, j})
			}

			if grid[i][j] > 1 {
				extraC.i = i
				extraC.j = j
			}
		}
	}

	if len(zerosC) == 0 {
		return 0
	}

	moves := make([]int, 0)

	for _, zeroC := range zerosC {
		newGrid := deepCopy(grid)

		newGrid[extraC.i][extraC.j] -= 1
		newGrid[zeroC.i][zeroC.j] += 1

		move := diff(extraC, zeroC)
		move += minimumMoves(newGrid)

		moves = append(moves, move)
	}

	return getMin(moves)
}

func getMin(arr []int) int {
	minInt := math.MaxInt

	for _, v := range arr {
		if v < minInt {
			minInt = v
		}
	}

	return minInt
}

func deepCopy(src [][]int) [][]int {
	dst := make([][]int, len(src))
	for i := range src {
		dst[i] = append([]int(nil), src[i]...)
	}

	return dst
}

func diff(a, b coordinate2) int {
	d := 0

	if a.i > b.i {
		d += a.i - b.i
	} else {
		d += b.i - a.i
	}

	if a.j > b.j {
		d += a.j - b.j
	} else {
		d += b.j - a.j
	}

	return d
}
