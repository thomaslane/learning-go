package main

import "fmt"

type Node struct {
	prev *Node
	next *Node
	key  interface{}
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

func (l *List) Display() {
	for list := l.head; list != nil; list = list.next {
		fmt.Printf("%+v ->", list.key)
	}
	fmt.Println()
}

func Display(n *Node) {
	for ; n != nil; n = n.next {
		fmt.Printf("%+v ->", n.key)
	}
	fmt.Println()
}

func ShowBackwards(n *Node) {
	for ; n != nil; n = n.prev {
		fmt.Printf("%+v ->", n.key)
	}
	fmt.Println()
}

func (l *List) Reverse() {
	for curr := l.head; curr != nil; curr = curr.prev {
		curr.next, curr.prev = curr.prev, curr.next
	}

	l.head, l.tail = l.tail, l.head

	Display(l.head)
}

func main() {
	link := List{}

	link.Insert(5)
	link.Insert(2)
	link.Insert(3)
	link.Insert(1)
	link.Insert(4)

	fmt.Println(link.head, link.tail)

	link.Display()

	ShowBackwards(link.tail)

	link.Reverse()

	fmt.Println(link.head, link.tail)
}
