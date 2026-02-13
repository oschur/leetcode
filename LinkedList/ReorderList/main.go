package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	tmp := premiddleList(head)
	p2 := reverseList(tmp.Next)

	tmp.Next = nil

	p1 := head
	for p2 != nil {
		p1NextTmp := p1.Next
		p1.Next = p2
		p1 = p2
		p2 = p1NextTmp

	}
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		tmp := cur
		cur = cur.Next
		tmp.Next = prev
		prev = tmp
	}
	return prev
}

func premiddleList(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
