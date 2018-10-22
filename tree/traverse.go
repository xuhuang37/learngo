package tree

func (node *Node)Traverse() {
	node.TraverseFuc(func(n *Node) {
		n.Print()
	})
}

func (node *Node)TraverseFuc(f func(node *Node)) {
	if node == nil {
		return
	}
	node.Left.TraverseFuc(f)
	f(node)
	node.Right.TraverseFuc(f)
}

func (node *Node)TraverseWithChannel() chan *Node {
	out:= make(chan *Node)
	go func() {
		node.TraverseFuc(func(node *Node) {
			out<-node
		})
		close(out)
	}()
	return out
}