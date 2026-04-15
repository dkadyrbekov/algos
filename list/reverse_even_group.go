package list

func reverseEvenLengthGroups(head *ListNode) {
	dumb := &ListNode{Next: head}
	fast, slow := dumb, dumb

	nextGroupSize := 1
	for fast.Next != nil {
		currentGroupSize := 0
		for {
			currentGroupSize++
			fast = fast.Next

			if fast.Next == nil {
				currentGroupSize++
				break
			}

			if currentGroupSize == nextGroupSize {
				break
			}
		}

		if currentGroupSize%2 == 0 {
			prev := fast.Next
			current := slow.Next
			next := current.Next

			for {
				current.Next = prev
				prev = current
				current = next
				next = next.Next

				if current == fast {
					current.Next = prev
					break
				}
			}

			slow.Next = fast
		}

		slow = fast
		nextGroupSize++
	}
}
