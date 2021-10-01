package main

import (
	"fmt"
)

func main() {
	var (
		node *Node
		err  error
	)
	node = node.AddFirst(1)
	node = node.AddFirst(0)
	node = node.AddFirst(-1)
	node = node.AddLast(2)
	node = node.AddLast(3)

	node, err = node.DeleteFirst()
	node, err = node.DeleteLast()
	node, err = node.DeleteFirst()
	node, err = node.DeleteLast()
	node, err = node.DeleteFirst()
	node, err = node.DeleteLast()
	node, err = node.DeleteFirst()

	fmt.Println(node, err)
}

// type for standard linked list

type Node struct {
	// the data in the linked list node
	Val int
	// pointer to the next node
	Next *Node
}

// return new head, after adding first

func (head *Node) AddFirst(e int) *Node {
	// create a new node
	node := &Node{Val: e}
	// current first now will become second
	node.Next = head
	// and return new node as head
	return node
}

// return new head, after adding at last

func (head *Node) AddLast(e int) *Node {
	// create a new node
	node := &Node{Val: e}

	if head == nil {
		return node
	}

	// iterate till last node
	iter := head
	for iter.Next != nil {
		iter = iter.Next
	}

	//add new node after last node
	iter.Next = node
	return head
}

// return new head, after deleting first

func (head *Node) DeleteFirst() (*Node, error) {
	if head == nil {
		return nil, fmt.Errorf("empty list")
	}
	// simply send whatever is there at 2nd pos
	return head.Next, nil
}

// return new head, after deleting last

func (head *Node) DeleteLast() (*Node, error) {

	if head == nil {
		return nil, fmt.Errorf("empty list")
	}
	// if head has only one element, then it would be last
	if head.Next == nil {
		return nil, nil
	}
	// reach till second last node
	for head.Next.Next != nil {
		head = head.Next
	}
	//then remove it's next, which is pointing to last element
	head.Next = nil
	return head, nil
}

// return last node of list

func (head *Node) Last() *Node {
	if head == nil {
		return nil
	}
	// iterate till last node
	for head.Next != nil {
		head = head.Next
	}
	return head
}

func (head *Node) Display() {
	// iterate till head has a Val
	for head != nil {
		fmt.Print(head, " > ")
		head = head.Next
	}
	fmt.Println()
}
