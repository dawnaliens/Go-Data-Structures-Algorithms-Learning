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

func (n *Node) Delete(k int) *Node {
	if n == nil {
		return nil
	}
	if k < n.Key {
		n.Left = n.Left.Delete(k)
		return n
	}
	if k > n.Key {
		n.Right = n.Right.Delete(k)
		return n
	}

	if n.Left == nil && n.Right == nil {
		return nil
	}

	if n.Left == nil {
		return n.Right
	}
	if n.Right == nil {
		return n.Left
	}
	minRight := n.Right.minNode()
	n.Key = minRight.Key
	n.Right = n.Right.Delete(n.Key)
	return n
}

// Find the smallest node in the tree
func (n *Node) minNode() *Node {
	current := n
	for current.Left != nil {
		current = current.Left
	}
	return current
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

	fmt.Println("In Order Traversal:")
	tree.InOrderTraversal()
	tree.Delete(1)
	fmt.Println("After Delete: ")
	tree.InOrderTraversal()
}
