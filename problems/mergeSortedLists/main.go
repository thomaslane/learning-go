// merge two sorted lists into one sorted list

package main

import "fmt"

type List struct {
	Head *Node
	Tail *Node
}

type Node struct {
	Val  int
	Next *Node
}

func (l *List) push(n *Node) {
	if l.Head == nil {
		l.Head = n
		l.Tail = n
	}

	l.Tail.Next = n
	l.Tail = n
}

func (l *List) display() {
	if l.Head == nil {
		return
	}

	for node := l.Head; node != nil; node = node.Next {
		fmt.Println(node.Val)
	}
}

func mergeSortedLists(listA, listB *List) *List {
	mergedList := &List{}
	var tail *Node
	tempA, tempB := listA.Head, listB.Head

	if listA == nil {
		return listB
	} else if tempB == nil {
		return listA
	}

	if tempA.Val < tempB.Val {
		mergedList.Head = &Node{Val: tempA.Val}
		tail = mergedList.Head
		tempA = tempA.Next
	} else {
		mergedList.Head = &Node{Val: tempB.Val}
		tail = mergedList.Head
		tempB = tempB.Next
	}

	for tempA != nil && tempB != nil {
		var temp Node
		if tempA.Val < tempB.Val {
			temp = *tempA
			tempA = tempA.Next
		} else {
			temp = *tempB
			tempB = tempB.Next
		}

		tail.Next = &temp
		tail = &temp
	}

	if tempA != nil {
		tail.Next = tempA
	} else if tempB != nil {
		tail.Next = tempB
	}

	return mergedList
}

func main() {
	listA, listB := &List{}, &List{}

	listA.push(&Node{Val: 3})
	listA.push(&Node{Val: 4})
	listA.push(&Node{Val: 6})
	listA.push(&Node{Val: 8})
	listA.push(&Node{Val: 9})
	listA.push(&Node{Val: 10})

	listB.push(&Node{Val: 1})
	listB.push(&Node{Val: 2})
	listB.push(&Node{Val: 7})

	mergedList := mergeSortedLists(listA, listB)

	mergedList.display()
	fmt.Println("-----")
	listA.display()
	fmt.Println("-----")
	listB.display()
}
