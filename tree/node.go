package tree

import "fmt"

type Node struct {
	Value int
	Left,Right *Node
}

func  (node Node)Print()  {
	fmt.Println(node.Value)
}
func (node *Node) SetValue(value int)  {
	if node == nil {
		fmt.Println("Setting Value to nil. Ignored")
		return
	}
	node.Value = value
}


func CreateNode(i int) *Node {
	return &Node{Value: i}
}

