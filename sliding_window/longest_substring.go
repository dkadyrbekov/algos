package sliding_window

func findLongestSubstring(str string) int {

	strArr := []rune(str)

	if len(strArr) == 0 {
		return 0
	}

	left, right := 0, 0
	maxLen := 1
	lettersMap := map[rune]int{
		strArr[0]: 1,
	}

	for right < len(strArr)-1 {
		right++
		lettersMap[strArr[right]]++

		for lettersMap[strArr[right]] > 1 {
			lettersMap[strArr[left]]--
			left++
		}

		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}

	return maxLen
}
