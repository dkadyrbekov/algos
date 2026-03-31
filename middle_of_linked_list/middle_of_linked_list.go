package middle_of_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func getMiddleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head

	for {
		if fast == nil || fast.Next == nil {
			return slow
		}

		slow = slow.Next
		fast = fast.Next.Next
	}
}
