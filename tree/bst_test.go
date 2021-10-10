package tree

import "testing"

func TestBST(t *testing.T) {
	var node *Node
	node = node.Add(10)
	node = node.Add(5)
	node = node.Add(15)
	node = node.Add(20)
	node = node.Add(8)
	node.Inorder()

	if node.Max() != nil && node.Max().Val != 20 {
		t.Errorf("max should be 20, is %v", node.Max())
	}
	if node.Min() != nil && node.Min().Val != 5 {
		t.Errorf("min should be 5, is %v", node.Min())
	}
	if node.MaxHeight() != 3 {
		t.Errorf("height should be 3, is %d", node.MaxHeight())
	}
}
