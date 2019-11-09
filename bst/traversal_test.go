package bst

import "testing"

func TestInOrderTraversalRecursive(t *testing.T) {
	initialBst := []int{5, 3, 7, 2, 4, 6, 8, 1}
	bst := initBst(initialBst)

	traversalResult := []int{}
	visitNode := func(k, d interface{}) {
		traversalResult = append(traversalResult, k.(int))
	}

	bst.InOrderRec(visitNode)

	expectedResult := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for i, expected := range expectedResult {
		if traversalResult[i] != expected {
			t.Errorf("expected traversalResult at index: %d to equal: %d instead got %d", i, expected, traversalResult[i])
		}
	}
}

func TestInOrderTraversalIterative(t *testing.T) {
	initialBst := []int{5, 3, 7, 2, 4, 6, 8, 1}
	bst := initBst(initialBst)

	traversalResult := []int{}
	visitNode := func(k, d interface{}) {
		if k == nil {
			return
		}
		traversalResult = append(traversalResult, k.(int))
	}

	bst.InOrder(visitNode)

	expectedResult := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for i, expected := range expectedResult {
		if traversalResult[i] != expected {
			t.Errorf("expected traversalResult at index: %d to equal: %d instead got %d", i, expected, traversalResult[i])
		}
	}
}

func TestPreOrderTraversalRecursive(t *testing.T) {
	initialBst := []int{5, 3, 7, 2, 4, 6, 8, 1}
	bst := initBst(initialBst)

	traversalResult := []int{}
	visitNode := func(k, d interface{}) {
		traversalResult = append(traversalResult, k.(int))
	}

	bst.PreOrderRec(visitNode)

	expectedResult := []int{5, 3, 2, 1, 4, 7, 6, 8}
	for i, expected := range expectedResult {
		if traversalResult[i] != expected {
			t.Errorf("expected traversalResult at index: %d to equal: %d instead got %d", i, expected, traversalResult[i])
		}
	}
}

func TestPreOrderTraversalIterative(t *testing.T) {
	initialBst := []int{5, 3, 7, 2, 4, 6, 8, 1}
	bst := initBst(initialBst)

	traversalResult := []int{}
	visitNode := func(k, d interface{}) {
		traversalResult = append(traversalResult, k.(int))
	}

	bst.PreOrder(visitNode)

	expectedResult := []int{5, 3, 2, 1, 4, 7, 6, 8}
	for i, expected := range expectedResult {
		if traversalResult[i] != expected {
			t.Errorf("expected traversalResult at index: %d to equal: %d instead got %d", i, expected, traversalResult[i])
		}
	}
}

func TestPostOrderTraversalRecursive(t *testing.T) {
	initialBst := []int{5, 3, 7, 2, 4, 6, 8, 1}
	bst := initBst(initialBst)

	traversalResult := []int{}
	visitNode := func(k, d interface{}) {
		traversalResult = append(traversalResult, k.(int))
	}

	bst.PostOrderRec(visitNode)

	expectedResult := []int{1, 2, 4, 3, 6, 8, 7, 5}
	for i, expected := range expectedResult {
		if traversalResult[i] != expected {
			t.Errorf("expected traversalResult at index: %d to equal: %d instead got %d", i, expected, traversalResult[i])
		}
	}
}

func TestPostOrderTraversalIterative(t *testing.T) {
	initialBst := []int{5, 3, 7, 2, 4, 6, 8, 1}
	bst := initBst(initialBst)

	traversalResult := []int{}
	visitNode := func(k, d interface{}) {
		traversalResult = append(traversalResult, k.(int))
	}

	bst.PostOrder(visitNode)

	expectedResult := []int{1, 2, 4, 3, 6, 8, 7, 5}
	for i, expected := range expectedResult {
		if traversalResult[i] != expected {
			t.Errorf("expected traversalResult at index: %d to equal: %d instead got %d", i, expected, traversalResult[i])
		}
	}
}

func TestLevelOrderTraversal(t *testing.T) {
	initialBst := []int{5, 3, 7, 2, 4, 6, 8, 1}
	bst := initBst(initialBst)

	traversalResult := []int{}
	visitNode := func(k, d interface{}) {
		traversalResult = append(traversalResult, k.(int))
	}

	bst.LevelOrder(visitNode)

	expectedResult := []int{5, 3, 7, 2, 4, 6, 8, 1}
	for i, expected := range expectedResult {
		if traversalResult[i] != expected {
			t.Errorf("expected traversalResult at index: %d to equal: %d instead got %d", i, expected, traversalResult[i])
		}
	}
}
