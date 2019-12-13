package heap

func New(t HeapType) *Heap {
	return &Heap{
		Data: make([]*HeapNode, 0, HEAP_DEFAULT_SIZE),
		t:    t,
	}
}

// Top returns top of heap, without removing it
// returns key, value, error
// O(1)
func (h *Heap) Top() (int, interface{}, error) {
	if len(h.Data) == 0 {
		return 0, nil, ErrEmptyHeap
	}

	min := *(h.Data[0])
	return min.Key, min.Value, nil
}

// ExtractTop returns top of heap element and removes it
// returns key, value, error
// O(Log N)
func (h *Heap) ExtractTop() (int, interface{}, error) {
	topKey, topVal, err := h.Top()
	if err != nil {
		return topKey, topVal, err
	}

	h.swapNodes(0, len(h.Data)-1)   // bring last node to front of heap
	h.Data = h.Data[:len(h.Data)-1] // create a new slice on existing array O(1) time

	// percolate element down
	parentIdx := 0
	for parentIdx < len(h.Data) {
		childIdx := h.getMinChildIdx(parentIdx)
		if childIdx != -1 && h.shouldSwapUp(childIdx, parentIdx) {
			h.swapNodes(childIdx, parentIdx)
			parentIdx = childIdx
		} else {
			break
		}
	}

	return topKey, topVal, nil
}

// Insert add key to heap
// O(Log N)
func (h *Heap) Insert(key int, value interface{}) {
	h.Data = append(h.Data, &HeapNode{key, value})

	// percolate element up
	childIdx := len(h.Data) - 1
	parentIdx := getParentIndex(childIdx)
	for childIdx > 0 && h.shouldSwapUp(childIdx, parentIdx) {
		h.swapNodes(childIdx, parentIdx)
		childIdx = parentIdx
		parentIdx = getParentIndex(childIdx)
	}
}
