package threesum

func threeSum(nums []int) [][]int {

	result := [][]int{}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					result = append(result, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}

	return unique(result)
}

func unique(zeroSums [][]int) [][]int {
	uniqueSums := [][]int{}
	uniqueMaps := []map[int]bool{}

	for _, sumArr := range zeroSums {
		sumIsUnique := true

		for _, uniqueMap := range uniqueMaps {
			if mapIsIdentique(uniqueMap, sumArr) {
				sumIsUnique = false
				break
			}
		}

		if sumIsUnique {
			newUniqueMap := map[int]bool{}
			for _, number := range sumArr {
				newUniqueMap[number] = true
			}

			uniqueMaps = append(uniqueMaps, newUniqueMap)
		}
	}

	for _, uniqueMap := range uniqueMaps {
		uniqueSum := []int{}

		for uniqueNumber, _ := range uniqueMap {
			uniqueSum = append(uniqueSum, uniqueNumber)
		}

		uniqueSums = append(uniqueSums, uniqueSum)
	}

	return uniqueSums
}

func mapIsIdentique(mapArr map[int]bool, arr []int) bool {
	for _, number := range arr {
		if !mapArr[number] {
			return false
		}
	}

	return true
}
