package threesum

func threeSum(nums []int) [][]int {
	var result [][]int

	if len(nums) < 3 {
		return result
	}

	nums = sort(nums)

	var lastNegativeIndex int
	var hasNegative bool
	var firstPositiveIndex int
	var hasPostive bool
	var zerocount int

	for i, val := range nums {
		if val < 0 {
			lastNegativeIndex = i
			hasNegative = true
		}
		if val > 0 {
			firstPositiveIndex = i
			hasPostive = true
			break
		}

		if val == 0 {
			zerocount++
		}
	}

	if zerocount >= 3 {
		result = append(result, []int{0, 0, 0})
	}

	if !hasPostive || !hasNegative {
		return result
	}

	positiveArr := nums[firstPositiveIndex:]
	notNegativeArr := nums[lastNegativeIndex+1:]
	negativeArr := nums[:lastNegativeIndex+1]

	result = append(result, search(negativeArr, positiveArr)...)
	result = append(result, search(notNegativeArr, negativeArr)...)

	return result
}

func search(croppedArr, searchArr []int) [][]int {
	var result [][]int

	for i := 0; i < len(croppedArr); i++ {
		if i > 0 && croppedArr[i] == croppedArr[i-1] {
			continue
		}

		for j := i + 1; j < len(croppedArr); j++ {
			if j > i+1 && croppedArr[j] == croppedArr[j-1] {
				continue
			}

			searchNumber := croppedArr[i] + croppedArr[j]

			if searchInArr(searchArr, -searchNumber) {
				result = append(result, []int{croppedArr[i], croppedArr[j], -searchNumber})
			}
		}
	}

	return result
}

func sort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 1; j < len(arr)-i; j++ {
			if arr[j-1] > arr[j] {
				tmp := arr[j-1]
				arr[j-1] = arr[j]
				arr[j] = tmp
			}
		}
	}

	return arr
}

func searchInArr(arr []int, searchNumber int) (found bool) {
	for _, num := range arr {
		if num == searchNumber {
			return true
		}
	}

	return false
}
