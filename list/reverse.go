package list

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	//fmt.Printf("head = %v left=%v, right=%v\n", printList(head), left, right)

	beforeLeft := &ListNode{Next: head}

	for i := 1; i < left; i++ {
		beforeLeft = beforeLeft.Next
	}
	leftPtr := beforeLeft.Next

	rightPtr := leftPtr
	for i := left; i < right; i++ {
		rightPtr = rightPtr.Next
	}
	afterRight := rightPtr.Next

	//fmt.Printf("head=%v\nbeforeLeft=%v\nleft=%v\nright=%v\nnext=%v\n", printList(head), printList(beforeLeft), printList(leftPtr), printList(rightPtr), printList(afterRight))

	prev := afterRight
	current := leftPtr
	next := current.Next
	for current != afterRight {
		current.Next = prev
		prev = current
		current = next
		if next != nil {
			next = next.Next
		}
	}

	if beforeLeft.Next == head {
		//fmt.Printf("reversedHead = %v\n", printList(rightPtr))
		return rightPtr
	}

	beforeLeft.Next = rightPtr

	//fmt.Printf("reversedHead = %v\n", printList(head))

	return head
}
