package bst

// BSTNode node construct for binary search tree
// Assumes only unique variables in BST
type BSTNode struct {
	Key   interface{}
	Data  interface{}
	Left  *BSTNode
	Right *BSTNode
}

// CompareKeys compare function for BST node
type CompareKeys func(a, b interface{}) bool

// BST binary search tree construct
type BST struct {
	Head       *BSTNode
	LessThan   CompareKeys
	isKeyEqual CompareKeys
}
