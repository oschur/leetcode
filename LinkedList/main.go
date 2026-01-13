package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseListIterative(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var newHead *ListNode

	for head != nil {
		nextHead := head.Next
		head.Next = newHead
		newHead = head
		head = nextHead
	}

	return newHead
}

func reverseListRecurcive(head *ListNode) *ListNode {
	newHead := head

	if head != nil && head.Next != nil {
		newHead = reverseListRecurcive(head.Next)
		head.Next.Next = head
		head.Next = nil
	}

	return newHead
}
