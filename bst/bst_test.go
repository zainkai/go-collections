package bst

import (
	"errors"
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

func initBst(nodes []int) *BST {
	bst := New(isLessInt, isEqualInt)
	for _, key := range nodes { // populate tree
		bst.Insert(key, nil)
	}
	return bst
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
	nodes := []int{2, 1, 3, 4}
	bst := initBst(nodes)

	for _, key := range nodes {
		if n, err := bst.SearchNode(key); err != nil || n.Key.(int) != key {
			t.Errorf("could not search for key: %d in binary search tree", key)
		}
	}
}

func TestFindMin(t *testing.T) {
	nodes := []int{2, 1, 3, 4}
	bst := initBst(nodes)

	if key, _ := bst.FindMin(); key.(int) != 1 {
		t.Errorf("couldnt not find min key: %d", key)
	}
}

func TestFindMax(t *testing.T) {
	nodes := []int{2, 1, 3, 4}
	bst := initBst(nodes)

	if key, _ := bst.FindMax(); key.(int) != 4 {
		t.Errorf("couldnt not find max key: %d", key)
	}
}

func TestDeleteHead(t *testing.T) {
	nodes := []int{5, 3, 7, 2, 4, 6, 8, 1}
	bst := initBst(nodes)

	if err := bst.Delete(5); err != nil {
		t.Errorf("couldnt delete key: %d", 5)
	}
	if bst.Head.Key.(int) != 6 {
		t.Errorf("head key was not replaced correctly")
	}
}

// delete single node from bst tests
func TestDeleteNode(t *testing.T) {
	initialBst := []int{5, 3, 7, 2, 4, 6, 8, 1}
	tests := []int{5, 3, 7, 2, 4, 6, 8, 1}

	for _, target := range tests {
		bst := initBst(initialBst)

		if err := bst.Delete(target); err != nil { // delete target from bst
			t.Errorf("couldnt delete key: %d", target)
		}
		if _, err := bst.Search(target); !errors.Is(err, ErrNotFoundNode) { // check if target node is not in bst
			t.Errorf("%d was not deleted from bst", target)
		}

		traversalResult := []int{}
		visitNode := func(k, d interface{}) {
			traversalResult = append(traversalResult, k.(int))
		}
		bst.LevelOrder(visitNode)
		if len(traversalResult) != len(initialBst)-1 {
			t.Errorf("node was not delted from initial bst")
		}

		for _, node := range traversalResult {
			if node == target {
				t.Errorf("target node %d was not deleted", target)
			}
		}
	}
}
