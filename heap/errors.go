package heap

import "errors"

var (
	ErrChildNoExist error = errors.New("Child does not exist")
	ErrCouldNotSwap error = errors.New("Could not swap nodes in heap")
	ErrEmptyHeap    error = errors.New("Heap is empty")
)
