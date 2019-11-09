package bst

import (
	"github.com/zainkai/go-collections/queue"
	"github.com/zainkai/go-collections/stack"
)

// InOrderRec iterate over BST recursively
// O(N) time, O(N) space
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

// InOrder iterate over BST iteratively
// O(N) time, O(N) space
func (tree *BST) InOrder(f func(key, data interface{})) {
	if tree.Head == nil {
		return
	}

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

// PreOrderRec iterate over BST recursively
// O(N) time, O(N) space
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

// PreOrder iterate over BST iteratively
// O(N) time, O(N) space
func (tree *BST) PreOrder(f func(key, data interface{})) {
	if tree.Head == nil {
		return
	}

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

// PreOrderRec iterate over BST iteratively
// O(N) time, O(N) space
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

// PostOrder iterate over BST iteratively
// O(N) time, O(2N) space
func (tree *BST) PostOrder(f func(key, data interface{})) {
	if tree.Head == nil {
		return
	}

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

// LevelOrder iterate over BST iteratively
// O(N) time, O(N) space
func (tree *BST) LevelOrder(f func(key, data interface{})) {
	if tree.Head == nil {
		return
	}

	queue := queue.New()
	queue.Enqueue(tree.Head)

	for queue.Len() > 0 {
		temp := queue.Dequeue().(*NodeBST)
		f(temp.Key, temp.Data)

		if temp.Left != nil {
			queue.Enqueue(temp.Left)
		}
		if temp.Right != nil {
			queue.Enqueue(temp.Right)
		}
	}
}
