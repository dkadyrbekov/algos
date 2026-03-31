package happy_number

func isHappy(num int) bool {
	if num <= 0 {
		return false
	}

	fast, slow := num, num

	for {
		if slow == 1 {
			return true
		}
		slow = countNextNumber(slow)

		if fast == 1 {
			return true
		}
		fast = countNextNumber(fast)

		if fast == 1 {
			return true
		}
		fast = countNextNumber(fast)

		if fast == slow {
			return false
		}
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
