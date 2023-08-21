package main

import (
	"errors"
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

type StackUsingLinkedList struct {
	Top *Node
}

func (s *StackUsingLinkedList) Push(value int) {
	newNode := &Node{Value: value}
	if s.Top != nil {
		newNode.Next = s.Top
	}
	s.Top = newNode
}

func (s *StackUsingLinkedList) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("Stack is empty")
	}
	item := s.Top.Value
	s.Top = s.Top.Next
	return item, nil
}

func (s *StackUsingLinkedList) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("Stack is empty")
	}
	return s.Top.Value, nil
}

func (s *StackUsingLinkedList) IsEmpty() bool {
	return s.Top == nil
}

func main() {
	stackLinkedList := &StackUsingLinkedList{}
	stackLinkedList.Push(10)
	stackLinkedList.Push(20)
	stackLinkedList.Push(30)
	fmt.Println(stackLinkedList.Peek())
	fmt.Println(stackLinkedList.Pop())
	fmt.Println(stackLinkedList.Peek())
	stackLinkedList.Pop()
	fmt.Println(stackLinkedList.IsEmpty())
	fmt.Println(stackLinkedList.Peek())
}
