package main

import (
	"errors"
	"fmt"
)

type StackUsingSlice struct {
	items []int
}

func (s *StackUsingSlice) Push(item int) {
	s.items = append(s.items, item)
}

func (s *StackUsingSlice) Pop() (int, error) {
	length := len(s.items)
	if length == 0 {
		return 0, errors.New("Stack is empty")
	}

	item := s.items[length-1]
	s.items = s.items[:length-1]
	return item, nil
}

func (s *StackUsingSlice) Peek() (int, error) {
	length := len(s.items)
	if length == 0 {
		return 0, errors.New("Stack is empty")
	}

	return s.items[length-1], nil
}

func (s *StackUsingSlice) IsEmpty() bool {
	return len(s.items) == 0
}

func main() {
	stackSlice := &StackUsingSlice{}
	stackSlice.Push(1)
	stackSlice.Push(2)
	stackSlice.Push(3)

	fmt.Println(stackSlice.Peek())
	fmt.Println(stackSlice.Pop())
	fmt.Println(stackSlice.Peek())
	stackSlice.Pop()
	fmt.Println(stackSlice.IsEmpty())
	fmt.Println(stackSlice.Peek())
}
