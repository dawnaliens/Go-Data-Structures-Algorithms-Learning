package main

import (
	"errors"
	"fmt"
)

type SliceQueue struct {
	array []int
}

func (s *SliceQueue) Enqueue(element int) {
	s.array = append(s.array, element)
}

func (s *SliceQueue) Dequeue() (int, error) {
	s.array = s.array[1:len(s.array)]
	if len(s.array) == 0 {
		return 0, errors.New("The queue is empty.")
	}
	return len(s.array), nil
}

func (s *SliceQueue) IsEmpty() bool {
	if len(s.array) == 0 {
		return true
	}
	return false
}

func (s *SliceQueue) First() (int, error) {
	if len(s.array) == 0 {
		return 0, errors.New("The Queue is empty")
	}
	return s.array[0], nil
}

func main() {
	Queue := &SliceQueue{}
	Queue.Enqueue(23)
	Queue.Enqueue(45)
	Queue.Enqueue(99)
	Queue.Enqueue(100)
	Queue.Enqueue(200)

	fmt.Println(Queue.IsEmpty())
	fmt.Println(Queue)
	fmt.Println(Queue.First())
	Queue.Dequeue()
	fmt.Println(Queue)
}
