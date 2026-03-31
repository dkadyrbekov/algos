package happy_number

func isHappy(num int) bool {
	if num <= 0 {
		return false
	}

	passedNumbers := make(map[int]bool)

	for {
		if num == 1 {
			return true
		}

		if passedNumbers[num] {
			return false
		}
		passedNumbers[num] = true

		num = countNextNumber(num)
	}
}

func countNextNumber(num int) int {
	if num <= 0 {
		return 0
	}

	nextNumber := 0
	intDevider := 1
	remainderDevider := 10

	for num/intDevider >= 1 {
		digit := (num % remainderDevider) / intDevider
		nextNumber += digit * digit

		intDevider *= 10
		remainderDevider *= 10
	}

	return nextNumber
}
