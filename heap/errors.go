package heap

import "errors"

var (
	ErrCouldNotSwap error = errors.New("Could not swap nodes in heap")
	ErrEmptyHeap    error = errors.New("Heap is empty")
)
