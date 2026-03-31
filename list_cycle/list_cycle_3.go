package list_cycle

func countCycleLength(head *ListNode) int {
	if head == nil {
		return 0
	}

	fast, slow := head, head

	for {
		slow = slow.Next
		if fast == nil || fast.Next == nil {
			return 0
		}
		fast = fast.Next.Next

		if slow == fast {
			break
		}
	}

	cycleLength := 1
	for {
		fast = fast.Next
		if fast == slow {
			return cycleLength
		}

		cycleLength++
	}
}
