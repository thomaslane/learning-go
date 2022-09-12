// given a string containing parenthesis characters, validate that it is a balanced expression of parentheses

package main

import "fmt"

type Node struct {
	Value interface{}
	Next  *Node
}

type Stack struct {
	Top    *Node
	Length int
}

func (s *Stack) push(v interface{}) {
	newNode := &Node{Value: v, Next: s.Top}
	s.Top = newNode
	s.Length++
}

func (s *Stack) pop() interface{} {
	n := s.Top
	if n != nil {
		s.Top = n.Next
		s.Length--
		return n.Value
	}
	return nil
}

func (s *Stack) getLength() int {
	return s.Length
}

func areBracketsBalanced(exp string) bool {
	s := &Stack{}

	for _, v := range exp {
		strV := string(v)
		switch strV {
		case "(", "{", "[":
			s.push(strV)
		case ")":
			check := s.pop()
			if check == nil || check == "{" || check == "[" {
				return false
			}
		case "}":
			check := s.pop()
			if check == nil || check == "(" || check == "[" {
				return false
			}
		case "]":
			check := s.pop()
			if check == nil || check == "(" || check == "{" {
				return false
			}
		}
	}

	return s.getLength() == 0
}

func main() {
	sample := "func main() { fmt.Println(s) }"
	// questions
	// what does a balanced expression of parentheses mean?
	// what types of parenthesis?

	balanced := areBracketsBalanced(sample)

	if balanced {
		fmt.Println("Strings are balanced")
	} else {
		fmt.Println("Strings are not balanced")
	}
}
