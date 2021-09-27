package main

import (
	"fmt"
)

func main() {
	node := &Node{Val: 0}
	node = node.Add(1)
	node = node.Add(2)
	node = node.Add(3)
	node.Display()
	node = node.Delete()
	node.Display()
	node = node.Delete()
	node.Display()
	fmt.Println(node.Len())
}

type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

func (head *Node) Add(element int) *Node {
	ele := &Node{Val: element}
	if head == nil {
		head = ele
		ele.Next = ele
		ele.Prev = ele
		return head
	}
	ele.Prev = head
	ele.Next = head.Next
	head.Next = ele

	return head
}

func (head *Node) Delete() *Node {
	if head == nil {
		return nil
	}
	curr = head.Next

	return head
}

func (head *Node) Display() {
	for head != nil {
		fmt.Printf("%v -> ", head.Val)
		head = head.Next
	}
	fmt.Println("")
}

func (head *Node) Len() int {
	l := 0
	for head != nil {
		l++
		head = head.Next
	}
	return l
}
