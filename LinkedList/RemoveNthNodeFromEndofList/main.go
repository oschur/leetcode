package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	current := dummy
	lenList := 1

	for current.Next != nil {
		lenList++
		current = current.Next
	}

	pos := lenList - n - 1
	current = dummy

	for i := 0; i < pos; i++ {
		current = current.Next
	}

	current.Next = current.Next.Next

	return dummy.Next
}
