package main

import (
	"fmt"
	"leango/tree"
)

type myTreeNode struct {
	*tree.Node //Embedding
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.Node == nil {
		return
	}
	left := myTreeNode{Node: myNode.Left}
	right := myTreeNode{Node: myNode.Right}
	left.postOrder()
	right.postOrder()
	myNode.Print()
}

func (myNode *myTreeNode) Traverse() { //类似重载
	fmt.Println("this method is shadowed")
}

func main() {
	root := myTreeNode{&tree.Node{Value: 3}}
	fmt.Println(root)

	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)

	root.Traverse()      //重载的函数
	root.Node.Traverse() //调用原本的函数

	fmt.Println()
	root.postOrder()
	fmt.Println()

	//var baseRoot *tree.Node
	// baseRoot := &root

	c := root.TraverseWithChannel()
	maxNode := 0
	for node := range c {
		if node.Value > maxNode {
			maxNode = node.Value
		}
	}
	fmt.Println("Max node value:", maxNode)
}
