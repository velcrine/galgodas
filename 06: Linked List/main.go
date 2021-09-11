package main

import (
	"fmt"
)

func main() {
	node := &Node{Val: 0}
	node = node.AddFront(1)
	node = node.AddFront(2)
	node = node.AddFront(3)
	node.AddLast(1)
	node.AddLast(2)
	node.AddLast(3)
	node.Display()
	node = node.DeleteFront()
	node.Display()
	node.DeleteLast()
	node.Display()
	fmt.Println(node.Len())
}

type Node struct {
	Val  int
	Next *Node
}

func (head *Node) AddFront(element int) *Node {
	ele := &Node{Val: element}
	if head == nil {
		head = ele
		return head
	}
	ele.Next = head
	return ele
}

func (head *Node) AddLast(element int) {
	ele := &Node{Val: element}
	if head == nil {
		head = ele
		return
	}
	for head.Next != nil {
		head = head.Next
	}
	head.Next = ele
}

func (head *Node) DeleteFront() *Node {
	if head == nil {
		return nil
	}
	head = head.Next
	return head
}

func (head *Node) DeleteLast() (element int) {
	if head == nil {
		return -1
	} else if head.Next == nil {
		temp := head.Val
		head = nil
		return temp
	}
	var prev *Node
	for head.Next != nil {
		prev = head
		head = head.Next
	}
	prev.Next = nil
	return head.Val
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

func removeDuplicate() {

}
