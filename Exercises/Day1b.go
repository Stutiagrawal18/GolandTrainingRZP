package main

import "fmt"

type Node struct {
	Value string
	Left  *Node
	Right *Node
}

func preOrder(n *Node) {
	if n == nil {
		return
	}
	fmt.Print(n.Value + " ")
	preOrder(n.Left)
	preOrder(n.Right)
}

func postOrder(n *Node) {
	if n == nil {
		return
	}
	fmt.Print(n.Value + " ")
	postOrder(n.Right)
	postOrder(n.Left)
}

func main() {

	a := &Node{Value: "a"}
	b := &Node{Value: "b"}
	c := &Node{Value: "c"}

	sub := &Node{Value: "-", Left: b, Right: c}
	root := &Node{Value: "+", Left: a, Right: sub}

	fmt.Print("PreOrder: ")
	preOrder(root)
	fmt.Println()

	fmt.Print("PostOrder: ")
	postOrder(root)
	fmt.Println()
}
