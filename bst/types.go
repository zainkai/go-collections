package bst

// NodeBST node construct for binary search tree
// Assumes only unique variables in BST
type NodeBST struct {
	Key   interface{}
	Data  interface{}
	Left  *NodeBST
	Right *NodeBST
}

// CompareKeys compare function for BST node
type CompareKeys func(a, b interface{}) bool

// BST binary search tree construct
type BST struct {
	Head       *NodeBST
	LessThan   CompareKeys
	isKeyEqual CompareKeys
}
