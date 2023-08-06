package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (l *LinkedList) Add(data int) {
	AnotherNode := &Node{Data: data}
	if l.Head == nil {
		l.Head = AnotherNode
		return
	}
	currentNode := l.Head
	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}
	currentNode.Next = AnotherNode
}

func (l *LinkedList) Delete(data int) {
	if l.Head == nil {
		return
	}
	if l.Head.Data == data {
		l.Head = l.Head.Next
		return
	}
	currentNode := l.Head
	for currentNode.Next != nil {
		if currentNode.Next.Data == data {
			currentNode.Next = currentNode.Next.Next
			return
		}
		currentNode = currentNode.Next
	}
}

func (l *LinkedList) Display() {
	currentNode := l.Head
	for currentNode != nil {
		fmt.Println(currentNode.Data)
		currentNode = currentNode.Next
	}
}

func main() {
	ll := &LinkedList{}
	ll.Add(1)
	ll.Add(2)
	ll.Add(3)
	ll.Add(4)

	ll.Display()
	fmt.Println("After deleting 3:")
	ll.Delete(3)
	ll.Display()
}
