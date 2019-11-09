package bst

import (
	"github.com/zainkai/go-collections/stack"
)

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

func (tree *BST) InOrder(f func(key, data interface{})) {
	stk := stack.New()
	temp := tree.Head

	for temp != nil || stk.Len() > 0 {
		if temp != nil {
			stk.Push(temp)
			temp = temp.Left
			continue
		}
		temp = stk.Pop().(*NodeBST)
		f(temp.Key, temp.Data)

		temp = temp.Right
	}
}

func (tree *BST) PreOrderRec(f func(key, data interface{})) {
	tree.Head.preOrderRec(f)
}

func (n *NodeBST) preOrderRec(f func(key, data interface{})) {
	if n == nil {
		return
	}

	f(n.Key, n.Data)
	n.Left.preOrderRec(f)
	n.Right.preOrderRec(f)
}

func (tree *BST) PreOrder(f func(key, data interface{})) {
	stk := stack.New()
	stk.Push(tree.Head)

	for stk.Len() > 0 {
		temp := stk.Pop().(*NodeBST)
		f(temp.Key, temp.Data)

		if temp.Right != nil {
			stk.Push(temp.Right)
		}
		if temp.Left != nil {
			stk.Push(temp.Left)
		}
	}
}

func (tree *BST) PostOrderRec(f func(key, data interface{})) {
	tree.Head.postOrderRec(f)
}

func (n *NodeBST) postOrderRec(f func(key, data interface{})) {
	if n == nil {
		return
	}

	n.Left.postOrderRec(f)
	n.Right.postOrderRec(f)
	f(n.Key, n.Data)
}

// PostOrder 2 stack solution
func (tree *BST) PostOrder(f func(key, data interface{})) {
	stk := stack.New()
	stk.Push(tree.Head)
	outputStk := stack.New()

	for stk.Len() > 0 {
		temp := stk.Pop().(*NodeBST)
		outputStk.Push(temp)

		if temp.Left != nil {
			stk.Push(temp.Left)
		}
		if temp.Right != nil {
			stk.Push(temp.Right)
		}
	}

	for outputStk.Len() > 0 {
		temp := outputStk.Pop().(*NodeBST)
		f(temp.Key, temp.Data)
	}
}
