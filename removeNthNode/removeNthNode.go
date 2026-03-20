package removeNthNode

// Definition for a Linked List node

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthLastNode(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	var listLength int

	for h := head; h != nil; h = h.Next {
		listLength++
	}

	nthFromHead := listLength - n

	// Некорретный ввод, ничего не удаляем
	if nthFromHead < 0 {
		return head
	}

	removeNode := head
	beforRemoveNode := head
	for i := 0; i < nthFromHead; i++ {
		removeNode = removeNode.Next

		if i != 0 {
			beforRemoveNode = beforRemoveNode.Next
		}
	}

	if removeNode == beforRemoveNode {
		return head.Next
	}

	beforRemoveNode.Next = removeNode.Next

	return head
}
