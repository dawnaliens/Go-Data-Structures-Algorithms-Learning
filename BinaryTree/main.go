package main

import "fmt"

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(k int) {
	if n == nil {
		return
	} else if k <= n.Key {
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	}
}

func (n *Node) InOrderTraversal() {
	if n == nil {
		return
	}
	n.Left.InOrderTraversal()
	fmt.Println(n.Key)
	n.Right.InOrderTraversal()
}

func (n *Node) PreOrderTraversal() {
	if n == nil {
		return
	}
	fmt.Println(n.Key)
	n.Left.PreOrderTraversal()
	n.Right.PreOrderTraversal()
}

func main() {
	tree := &Node{Key: 100}
	tree.Insert(200)
	tree.Insert(300)
	tree.Insert(400)
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(30)
	tree.Insert(25)
	tree.Insert(8)
	tree.Insert(5)
	fmt.Println("InOrderTraversal:")
	tree.PreOrderTraversal()
	fmt.Println("In Order Traversal:")
	tree.InOrderTraversal()
}
