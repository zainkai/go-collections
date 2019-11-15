package heap

func New(t HeapType) *Heap {
	return &Heap{
		Data: make([]*HeapNode, 0, HEAP_DEFAULT_SIZE),
		t:    t,
	}
}

func getParentIndex(index int) int {
	return (index - 1) / 2
}

func getLeftChildIndex(parentIndex int) int {
	return (2 * parentIndex) + 1
}

func getRightChildIndex(parentIndex int) int {
	return (2 * parentIndex) + 2
}

func (h *Heap) getKey(index int) int {
	return h.Data[index].key
}

func (h *Heap) swapNodes(i, j int) error {
	if i == j {
		return nil
	} else if i < 0 || j < 0 || i > len(h.Data)-1 || j > len(h.Data)-1 {
		return ErrCouldNotSwap
	}

	h.Data[i], h.Data[j] = h.Data[j], h.Data[i]
	return nil
}

// Top returns top of heap, without removing it
// O(1)
func (h *Heap) Top() (int, interface{}, error) {
	if len(h.Data) == 0 {
		return 0, nil, ErrEmptyHeap
	}

	min := *(h.Data[0])
	return min.key, min.value, nil
}

// ExtractTop returns top of heap element and removes it
// O(N)
func (h *Heap) ExtractTop() (int, interface{}, error) {
	topKey, topVal, err := h.Top()
	if err != nil {
		return topKey, topVal, err
	}

	return topKey, topVal, nil
}

// Insert add key to heap
// O(Log N)
func (h *Heap) Insert(key int, value interface{}) {
	h.Data = append(h.Data, &HeapNode{key, value})

	childIdx := len(h.Data) - 1
	parentIdx := getParentIndex(childIdx)
	for childIdx > 0 {
		childKey := h.getKey(childIdx)
		parentKey := h.getKey(parentIdx)
		if (h.t == MIN && childKey < parentKey) || (h.t == MAX && childKey > parentKey) {
			h.swapNodes(childIdx, parentIdx)
			childIdx = parentIdx
			parentIdx = getParentIndex(childIdx)
		} else {
			return
		}
	}
}
