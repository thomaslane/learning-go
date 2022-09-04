package main

import (
	"errors"
	"fmt"
	"sync"
)

// Stack struct
type Stack struct {
	Top    *Node
	Length int
	Sync   sync.Mutex
}

// Node struct
type Node struct {
	Value interface{}
	Next  *Node
}

// CreateStack method
func CreateStack() *Stack {
	return &Stack{}
}

// Push method
func (s *Stack) Push(v interface{}) {
	s.Sync.Lock()
	defer s.Sync.Unlock()

	newNode := &Node{Value: v, Next: s.Top}
	s.Top = newNode
	s.Length++
}

// Pop method
func (s *Stack) Pop() (*Node, error) {
	if s.Top != nil {
		s.Sync.Lock()
		defer s.Sync.Unlock()

		head := s.Top
		newHead := s.Top.Next
		s.Top = newHead
		s.Length--

		return head, nil
	} else {
		err := errors.New("cannot Pop, stack is empty")
		return nil, err
	}
}

// Display method
func (s *Stack) DisplayAll() error {
	if s.Top != nil {
		s.Sync.Lock()
		defer s.Sync.Unlock()

		for n := s.Top; n != nil; n = n.Next {
			fmt.Println(n.Value)
		}

		return nil
	} else {
		err := errors.New("cannot DisplayAll, stack is empty")
		return err
	}
}

// Peek
func (s *Stack) Peek() (*Node, error) {
	if s.Top != nil {
		return s.Top, nil
	} else {
		err := errors.New("cannot Peek, stack is empty")
		return nil, err
	}
}

func main() {
	s := CreateStack()
	fmt.Println(s)
	s.Push("First")
	fmt.Println(s)
	fmt.Println(s.Peek())
	fmt.Println(s.Pop())
	fmt.Println(s)
	fmt.Println(s.Pop())
	fmt.Println(s.Peek())
	s.Push("First")
	s.Push("Second")
	s.Push("Third")
	fmt.Println(s)
	s.DisplayAll()
	fmt.Println(s.Pop())
	s.DisplayAll()
	fmt.Println(s)
}
