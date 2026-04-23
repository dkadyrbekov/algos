package subsets

func letterCombinations(digits string) []string {
	return getCombinations("", []rune(digits))
}

func getCombinations(prefix string, digits []rune) []string {
	var digitLetters = map[rune][]rune{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}

	if len(digits) == 0 {
		return []string{prefix}
	}

	combinations := make([]string, 0)

	for _, dl := range digitLetters[digits[0]] {
		combinations = append(combinations, getCombinations(prefix+string(dl), digits[1:])...)
	}

	return combinations
}
