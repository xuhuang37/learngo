package main

import (
	"learngo/tree"
	"fmt"
)

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode)postOder()  {
	if myNode == nil || myNode.node == nil{
		return
	}
	left := myTreeNode{myNode.node.Left}
	left.postOder()
	right := myTreeNode{myNode.node.Right}
	right.postOder()
	myNode.node.Print()
}

func main() {
	var root tree.Node
	root = tree.Node{Value:3}
	root.Left = &tree.Node{Value:0}
	root.Left.Right = tree.CreateNode(2)
	root.Right = tree.CreateNode(5)
	root.Right.Left = tree.CreateNode(0)
	root.Traverse()
	fmt.Println("\n")
	mRoot:= myTreeNode{&root}
	mRoot.postOder()

}