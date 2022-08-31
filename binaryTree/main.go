package main

import "fmt"

var count int

// Node represents the components of a binary search tree
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// Insert will add a node to the tree
// they key to add should not already be in the tree
func (n *Node) Insert(k int) {
	if n.Key < k {
		// move right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		// move left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

// Search will take in a key value
// and RETURN true if there is a node with that value
func (n *Node) Search(k int) bool {
	count++
	if n == nil {
		return false
	} else if n.Key < k {
		// move right
		return n.Right.Search(k)
	} else if n.Key > k {
		// move left
		return n.Left.Search(k)
	}

	return true
}

func main() {
	tree := &Node{Key: 100}
	tree.Insert(52)
	fmt.Println(tree)
	tree.Insert(21)
	fmt.Println(tree)
	tree.Insert(45)
	fmt.Println(tree)
	tree.Insert(101)
	fmt.Println(tree)

	fmt.Println(tree.Search(21))
	fmt.Println(count)
}
