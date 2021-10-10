package tree

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func (n *Node) Add(e int) (node *Node) {
	node = &Node{Val: e}
	if n == nil {
		return
	}
	if n.Val > e {
		n.Left = n.Left.Add(e)
	} else if n.Val < e {
		n.Right = n.Right.Add(e)
	}
	return n
}

func (n *Node) Inorder() {
	if n == nil {
		return
	}
	n.Left.Inorder()
	fmt.Print(n.Val)
	n.Right.Inorder()
}

func (n *Node) Max() *Node {
	for n.Right != nil {
		n = n.Right
	}
	return n
}

func (n *Node) Min() *Node {
	for n.Left != nil {
		n = n.Left
	}
	return n
}

func (n *Node) MaxHeight() int32 {
	if n == nil {
		return 0
	}
	hl := n.Left.MaxHeight()
	hr := n.Right.MaxHeight()
	if hl > hr {
		return hl + 1
	}
	return hr + 1
}
