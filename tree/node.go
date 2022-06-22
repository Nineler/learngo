package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func (node Node) Print() {
	fmt.Print(node.Value)
}

func (node *Node) Traverse() {
	if node == nil {
		return
	}
	node.Print()
	node.Left.Traverse()
	node.Right.Traverse()
}

func CreateNode(value int) *Node {
	return &Node{Value: value}
}

func (node *Node) SetValue(value int) {
	if node == nil {
		return
	}
	node.Value = value
}
