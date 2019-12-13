package heap

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
	return h.Data[index].Key
}

// swapNodes swaps nodes at an index
func (h *Heap) swapNodes(i, j int) error {
	if i == j {
		return nil
	} else if i < 0 || j < 0 || i > len(h.Data)-1 || j > len(h.Data)-1 {
		return ErrCouldNotSwap
	}

	h.Data[i], h.Data[j] = h.Data[j], h.Data[i]
	return nil
}

// shouldSwapUp returns whether the child should replace its parent
func (h *Heap) shouldSwapUp(childIdx, parentIdx int) bool {
	if h.t == MIN {
		return h.getKey(childIdx) < h.getKey(parentIdx)
	}
	return h.getKey(childIdx) > h.getKey(parentIdx)
}

func (h *Heap) getMinChildIdx(parentIdx int) int {
	leftIdx := getLeftChildIndex(parentIdx)
	rightIdx := getRightChildIndex(parentIdx)
	if rightIdx >= len(h.Data) && leftIdx >= len(h.Data) {
		return -1
	} else if leftIdx >= len(h.Data) {
		return rightIdx
	} else if rightIdx >= len(h.Data) {
		return leftIdx
	}

	if h.shouldSwapUp(leftIdx, rightIdx) {
		return leftIdx
	}
	return rightIdx
}
