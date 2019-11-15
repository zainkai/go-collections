package heap

import "testing"

func TestMinHeapInsert(t *testing.T) {
	tests := []int{5, 4, 3, 2, 1, 0}
	heap := New(MIN)

	for _, testVal := range tests {
		heap.Insert(testVal, nil)

		top := *heap.Data[0]
		if top.key != testVal {
			t.Errorf("Inserting key %d did not rebalance to top of heap", testVal)
			t.Error(heap.Data)
		}
	}
}

func TestMaxHeapInsert(t *testing.T) {
	tests := []int{0, 1, 2, 3, 4, 5}
	heap := New(MAX)

	for _, testVal := range tests {
		heap.Insert(testVal, nil)

		top := *heap.Data[0]
		if top.key != testVal {
			t.Errorf("Inserting key %d did not rebalance to top of heap", testVal)
			t.Error(heap.Data)
		}
	}
}
