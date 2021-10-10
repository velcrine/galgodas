package main

/*
this node is defined in file linked-list/circular.go
type Node struct {
	Val  int
	Next *Node
}
*/

func (n *Node) Add(e int) (node *Node) {
	node = &Node{Val: e}
	if n == nil {
		n.Next = node
		return
	}
	node.Next = n
	iter := n
	for iter.Next != n {
		iter = iter.Next
	}
	iter.Next = node
	return n
}

func (n *Node) Delete() (node *Node) {
	if n == nil || n.Next == nil {
		return
	}
	iter := n
	prev := n
	for {
		iter = iter.Next

	}
	iter.Next = node
	return n
}
