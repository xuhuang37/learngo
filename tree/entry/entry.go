package main

import "learngo/tree"

func main() {
	var root tree.TreeNode
	root = tree.TreeNode{Value:3}
	root.Left = &tree.TreeNode{Value:0}
	root.Left.Right = tree.CreateNode(2)
	root.Right = tree.CreateNode(5)
	root.Right.Left = tree.CreateNode(0)
	root.Traverse()
}