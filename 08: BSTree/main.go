package main

import (
	"fmt"
)

func main() {
	node := &Node{
		Val: 18,
	}
	node = node.Insert(5)
	node = node.Insert(11)
	node = node.Insert(17)
	node = node.Insert(10)

	fmt.Println()
	node.PostOrderTraversal()
	fmt.Println(node.Height())

	fmt.Println(node.Max())
	fmt.Println(node.Min())
	fmt.Println(node.Search(17))
	fmt.Println(node.Search(19))
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

/*
  value = 7
  10 (L:sub-tree, R:nil)


*/

func (n *Node) Insert(v int) *Node {
	if n == nil {
		return &Node{
			Val: v,
		}
	}
	if v < n.Val {
		if n.Left == nil {
			n.Left = &Node{
				Val: v,
			}
		} else {
			n.Left.Insert(v)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{
				Val: v,
			}
		} else {
			n.Right.Insert(v)
		}
	}
	return n
}

func (n *Node) PostOrderTraversal() {

	if n == nil {
		return
	}
	n.Left.PostOrderTraversal()
	fmt.Print(n.Val, "  ")
	n.Right.PostOrderTraversal()
}

func (n *Node) Height() int {
	if n == nil {
		return 0
	}
	left := n.Left.Height()
	right := n.Right.Height()
	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}

func (n *Node) Max() int {

	if n.Right == nil {
		return n.Val
	} else {
		return n.Right.Max()
	}
}

func (n *Node) Min() int {

	if n.Left == nil {
		return n.Val
	} else {
		return n.Left.Min()
	}
}

func (n *Node) Search(ele int) *Node {
	if n == nil {
		return nil
	}
	if ele == n.Val {
		return n
	}
	if ele > n.Val {
		return n.Right.Search(ele)
	}
	if ele < n.Val {
		return n.Left.Search(ele)
	}
	return nil
}

func (n *Node) Delete(ele int) (tree *Node) {
	if n == nil {
		return n
	}
	if n.Val < ele {
		n.Left = n.Left.Delete(ele)
	} else if n.Val > ele {
		n.Right = n.Right.Delete(ele)
	} else {

		if n.Right == nil && n.Left == nil {
			return nil
		}
		if n.Right == nil {
			return n.Left
		}
		if n.Left == nil {
			return n.Right
		}
		m := n.Right.Min()
		n.Right = n.Right.Delete(m)
		n.Val = m
	}
	return n
}
