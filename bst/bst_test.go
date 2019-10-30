package bst

import (
	"testing"
)

var keysPreOrder = []int{4, 2, 6, 1, 3, 5, 7, 8}

func populatedBst() *BST {
	return nil
}

var isEqualInt CompareKeys = func(a, b interface{}) bool {
	return a.(int) == b.(int)
}

var isLessInt CompareKeys = func(a, b interface{}) bool {
	return a.(int) < b.(int)
}

func TestInsertBst(t *testing.T) {
	bst := New(isLessInt, isEqualInt)
	// bst.Insert(4, nil)

	if err := bst.Insert(2, nil); err != nil {
		t.Errorf("problem inserting into bst")
		t.Log(err)
	}
	if bst.Head == nil || !isEqualInt(bst.Head.Key, 2) {
		t.Errorf("head node did not have key 2")
	}

	if err := bst.Insert(1, nil); err != nil {
		t.Errorf("problem inserting into bst")
		t.Log(err)
	}
	if bst.Head.Left == nil || !isEqualInt(bst.Head.Left.Key, 1) {
		t.Errorf("head.Left node did not have key 1")
	}

	if err := bst.Insert(3, nil); err != nil {
		t.Errorf("problem inserting into bst")
		t.Log(err)
	}
	if bst.Head.Right == nil || !isEqualInt(bst.Head.Right.Key, 3) {
		t.Errorf("head.Left node did not have key 3")
	}

	if err := bst.Insert(4, nil); err != nil {
		t.Errorf("problem inserting into bst")
		t.Log(err)
	}
	if bst.Head.Right.Right == nil || !isEqualInt(bst.Head.Right.Right.Key, 4) {
		t.Errorf("head.Left node did not have key 4")
	}
}

func TestSearchBst(t *testing.T) {
	bst := New(isLessInt, isEqualInt)
	nodes := []int{2, 1, 3, 4}
	for _, key := range nodes { // populate tree
		bst.Insert(key, nil)
	}

	for _, key := range nodes {
		if n, err := bst.SearchNode(key); err != nil || n.Key.(int) != key {
			t.Errorf("could not search for key: %d in binary search tree", key)
		}
	}
}
