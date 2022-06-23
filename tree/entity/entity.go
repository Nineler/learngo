package main

import (
	"GOProject/tree"
	"fmt"
)

func main() {
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Right = &tree.Node{}
	root.Left = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)

	root.Traverse()
	nodecontent := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodecontent++
	})
	fmt.Println(nodecontent)
	c := root.TraverseWithChannel()
	maxNode := 0
	for node := range c {
		if node.Value > maxNode {
			maxNode = node.Value
		}
	}
	fmt.Printf("MaxValue:%d", maxNode)
}
