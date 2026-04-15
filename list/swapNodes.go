package list

func swapNodes(head *ListNode, k int) *ListNode {

	dumb := &ListNode{Next: head}

	befL := dumb
	for i := 0; i < k-1; i++ {
		befL = befL.Next
	}

	left := befL.Next

	befR := dumb
	curr := left

	for curr.Next != nil {
		befR = befR.Next
		curr = curr.Next
	}

	rigth := befR.Next

	befR.Next = left
	befL.Next = rigth
	tmp := rigth.Next
	rigth.Next = left.Next
	left.Next = tmp

	return dumb.Next
}
