package repeating_char

func longestRepeatingCharacterReplacement(s string, k int) int {
	runeStr := []rune(s)
	if len(runeStr) == 0 {
		return 0
	}

	maxLength := 1
	runesMap := map[rune]int{
		runeStr[0]: 1,
	}
	maxRuneCount := 1

	left := 0
	right := 0

	for right < len(runeStr)-1 {
		right++
		runesMap[runeStr[right]]++
		for _, v := range runesMap {
			if v > maxRuneCount {
				maxRuneCount = v
			}
		}

		if len(runeStr[left:right+1])-maxRuneCount > k {
			runesMap[runeStr[left]]--

			left++
		}

		if right-left+1 > maxLength {
			maxLength = right - left + 1
		}
	}

	return maxLength
}
