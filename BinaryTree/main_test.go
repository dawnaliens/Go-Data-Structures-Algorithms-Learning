package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	tree := &Node{Key: 50}
	tree.Insert(30)
	tree.Insert(70)
	tree.Insert(20)
	tree.Insert(40)
	tree.Insert(60)
	tree.Insert(80)

	buffer := new(bytes.Buffer)
	tree.InOrderTraversalWithBuffer(buffer)

	expected := "20\n30\n40\n50\n60\n70\n80\n"
	if expected != buffer.String() {
		t.Errorf("Expected %s but got %s", expected, buffer.String())
	}
}

func TestInOrderTraversal(t *testing.T) {
	tree := &Node{Key: 50}
	tree.Insert(30)
	tree.Insert(70)
	tree.Insert(20)
	tree.Insert(40)
	tree.Insert(60)
	tree.Insert(80)

	buffer := new(bytes.Buffer)
	tree.InOrderTraversalWithBuffer(buffer)

	expected := "20\n30\n40\n50\n60\n70\n80\n"
	if expected != buffer.String() {
		t.Errorf("Expected %s but got %s", expected, buffer.String())
	}
}

func TestPreOrderTraversal(t *testing.T) {
	tree := &Node{Key: 50}
	tree.Insert(30)
	tree.Insert(70)
	tree.Insert(20)
	tree.Insert(40)
	tree.Insert(60)
	tree.Insert(80)

	buffer := new(bytes.Buffer)
	tree.PreOrderTraversalWithBuffer(buffer)

	expected := "50\n30\n20\n40\n70\n60\n80\n"
	if expected != buffer.String() {
		t.Errorf("Expected %s but got %s", expected, buffer.String())
	}
}

func (n *Node) InOrderTraversalWithBuffer(buffer *bytes.Buffer) {
	if n == nil {
		return
	}
	n.Left.InOrderTraversalWithBuffer(buffer)
	buffer.WriteString(fmt.Sprintf("%d\n", n.Key))
	n.Right.InOrderTraversalWithBuffer(buffer)
}

func (n *Node) PreOrderTraversalWithBuffer(buffer *bytes.Buffer) {
	if n == nil {
		return
	}
	buffer.WriteString(fmt.Sprintf("%d\n", n.Key))
	n.Left.PreOrderTraversalWithBuffer(buffer)
	n.Right.PreOrderTraversalWithBuffer(buffer)
}
