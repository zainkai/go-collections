package bst

import "errors"

var (
	// ErrDuplicateNode node already exists in binary search tree and will not be re-added
	ErrDuplicateNode = errors.New("Duplicate node in BST")
	// ErrNotFoundNode not could not be found in binary search tree
	ErrNotFoundNode = errors.New("Node Not Found in BST")
)
