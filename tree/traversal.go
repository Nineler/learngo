package tree

import "fmt"

func (node *Node) Traverse() {
	node.TraverseFunc(func(n *Node) {
		n.Print()
	})
	fmt.Println()
}
func (node *Node) TraverseFunc(f func(*Node)) {
	if node == nil {
		return
	}
	node.Right.TraverseFunc(f)
	f(node)
	node.Left.TraverseFunc(f)
}

func (node *Node) TraverseWithChannel() chan *Node {
	out := make(chan *Node)
	go func() {
		node.TraverseFunc(func(node *Node) {
			out <- node
		})
		close(out)
	}()
	return out
}
