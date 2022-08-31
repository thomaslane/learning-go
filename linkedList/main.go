package main

import "fmt"

type Node struct {
	key  interface{}
	prev *Node
	next *Node
}

type List struct {
	head *Node
	tail *Node
}

func (L *List) Insert(key interface{}) {
	list := &Node{
		next: L.head,
		key:  key,
	}
	if L.head != nil {
		L.head.prev = list
	}
	L.head = list

	l := L.head
	for l.next != nil {
		l = l.next
	}
	L.tail = l
}

func (L *List) InsertAt(pos int, key interface{}) {
	newNode := &Node{
		key: key,
	}
	if pos < 0 {
		return
	}

	if pos == 0 {
		newNode.next = L.head
		L.head = newNode
	}

	cur := L.getAt(pos)
	prev := cur.prev

	cur.prev, prev.next = newNode, newNode

}

func (l *List) getAt(pos int) *Node {
	ptr := l.head
	if pos < 0 {
		return ptr
	}
	for i := 0; i < pos && ptr != nil; i++ {
		ptr = ptr.next
	}

	return ptr
}

func (l *List) Display() {
	list := l.head
	for list != nil {
		fmt.Printf("%v ->", list.key)
		list = list.next
	}
	fmt.Println()
}

func (L *List) Reverse() {
	var prev *Node
	cur := L.head
	for cur != nil {
		cur.next, cur.prev = cur.prev, cur.next
		prev = cur
		cur = cur.prev
	}
	L.head = prev
}

func main() {
	fmt.Println("Linked List")
	ll := &List{}

	ll.Insert("a")
	ll.Insert("b")
	ll.Insert("c")

	ll.Display()
	ll.Reverse()
	ll.Display()

	fmt.Println(ll.head, ll.tail)
}
