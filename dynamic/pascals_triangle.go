package dynamic

func generate(numRows int) [][]int {
	result := [][]int{{1}}

	for i := 2; i <= numRows; i++ {
		newArr := make([]int, 0, i)
		prevArr := result[i-2]
		for j := 0; j < i; j++ {
			newElem := 0

			if j-1 >= 0 {
				newElem += prevArr[j-1]
			}
			if j < len(prevArr) {
				newElem += prevArr[j]
			}

			newArr = append(newArr, newElem)
		}

		result = append(result, newArr)
	}

	return result
}
