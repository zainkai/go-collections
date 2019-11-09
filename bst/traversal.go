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

// func (tree *BST) InOrder(f func(key, data interface{})) {
// 	stk := stack.New()
// 	stk.Push(tree.Head)

// 	for stk.Len() > 0 {
// 		tmp := stk.Peek().(*NodeBST)
// 		if tmp.Left != nil {
// 			stk.Push(tmp.Left)
// 			continue
// 		} else {
// 			f(tmp.Key, tmp.Data)
// 			stk.Pop()

// 			if tmp.Right != nil {
// 				stk.Push(tmp.Right)
// 			}
// 		}

// 	}
// }
