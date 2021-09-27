package main

import (
	"fmt"
)

func main() {
	node := &Node{Val: 0}
	node = node.AddFront(1)
	node = node.AddFront(2)
	node = node.AddFront(3)
	node = node.AddLast(1)
	node = node.AddLast(2)
	node = node.AddLast(3)
	node.Display()
	node = node.DeleteFront()
	node.Display()
	node = node.DeleteLast()
	node.Display()
	fmt.Println(node.Len())
}

type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

// [87]
// [1] -> [2] -> [-4]

// head = [1]

// [87]<- head

// [87] ->(head) [1] -> [2] -> [-4]

// head = node[97]
// (head) [87] -> [1] -> [2] -> [-4]
// output
// [87] -> [1] -> [2] -> [-4]

func (head *Node) AddFront(element int) *Node {
	ele := &Node{Val: element}
	if head == nil {
		head = ele
		return head
	}

	ele.Next = head
	head.Prev = ele
	return ele
}

func (head *Node) AddLast(element int) *Node {
	ele := &Node{Val: element}
	curr := head
	if curr == nil {
		curr = ele
		return curr
	}
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = ele
	ele.Prev = curr

	return head
}

func (head *Node) DeleteFront() *Node {
	if head == nil {
		return nil
	}
	head = head.Next
	return head
}

func (head *Node) DeleteLast() *Node {
	curr := head
	if curr == nil {
		return nil
	} else if curr.Next == nil {
		return nil
	}
	var prev *Node
	for curr.Next != nil {
		prev = curr
		curr = curr.Next
	}
	prev.Next = nil
	return head
}

func (head *Node) Len() int {
	l := 0
	for head != nil {
		l++
		head = head.Next
	}
	return l
}

func (head *Node) Display() {
	for head != nil {
		fmt.Printf("%v -> ", head.Val)
		head = head.Next
	}
	fmt.Println("")
}
