package main

type LRUCache struct {
	Head, Tail *Node
	Cashe      map[int]*Node
	Capacity   int
}

type Node struct {
	Prev  *Node
	Next  *Node
	Key   int
	Value int
}

func newNode(key, val int) *Node {
	return &Node{
		Key:   key,
		Value: val,
	}
}

func Constructor(capacity int) LRUCache {
	head, tail := newNode(0, 0), newNode(0, 0)
	head.Next = tail
	tail.Prev = head
	return LRUCache{
		Head:     head,
		Tail:     tail,
		Cashe:    make(map[int]*Node),
		Capacity: capacity,
	}
}

func (this *LRUCache) Delete(node *Node) {
	delete(this.Cashe, node.Key)
	node.Prev.Next, node.Next.Prev = node.Next, node.Prev
}

func (this *LRUCache) Insert(node *Node) {
	this.Cashe[node.Key] = node
	next := this.Head.Next
	this.Head.Next = node
	node.Prev, next.Prev, node.Next = this.Head, node, next
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Cashe[key]; ok {
		this.Delete(node)
		this.Insert(node)
		return node.Value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.Cashe[key]; ok {
		this.Delete(this.Cashe[key])
	}

	if len(this.Cashe) == this.Capacity {
		this.Delete(this.Tail.Prev)
	}

	this.Insert(newNode(key, value))
}
