package bst

func (tree *BST) InOrderRec(f func(key, data interface{})) {
	tree.Head.inOrderRec(f)
}

func (n *NodeBST) inOrderRec(f func(key, data interface{})) {
	if n == nil {
		return
	}

	n.Left.inOrderRec(f)
	f(n.Key, n.Data)
	n.Right.inOrderRec(f)
}
