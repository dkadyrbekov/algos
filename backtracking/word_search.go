package backtracking

type coordinate struct {
	i int
	j int
}

func (c coordinate) equal(c2 coordinate) bool {
	return c.i == c2.i && c.j == c2.j
}

func wordSearch(grid [][]string, word string) bool {

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if findWord(grid, word, coordinate{i, j}, coordinate{-1, -1}) {
				return true
			}
		}
	}

	return false
}

func findWord(grid [][]string, word string, curr, prev coordinate) bool {
	if len(word) == 0 {
		return true
	}

	if grid[curr.i][curr.j][0] != word[0] {
		return false
	}

	if len(word) == 1 {
		return true
	}

	if curr.i > 0 {
		newC := curr
		newC.i--
		if !newC.equal(prev) && findWord(grid, word[1:], newC, curr) {
			return true
		}
	}
	if curr.i < len(grid)-1 {
		newC := curr
		newC.i++
		if !newC.equal(prev) && findWord(grid, word[1:], newC, curr) {
			return true
		}
	}
	if curr.j > 0 {
		newC := curr
		newC.j--
		if !newC.equal(prev) && findWord(grid, word[1:], newC, curr) {
			return true
		}
	}
	if curr.j < len(grid[curr.i])-1 {
		newC := curr
		newC.j++
		if !newC.equal(prev) && findWord(grid, word[1:], newC, curr) {
			return true
		}
	}

	return false
}
