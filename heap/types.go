package heap

type HeapNode struct {
	Key   int
	Value interface{}
}

type HeapType string

const (
	MIN               HeapType = "MIN"
	MAX               HeapType = "MAX"
	HEAP_DEFAULT_SIZE int      = 10
)

type Heap struct {
	Data []*HeapNode
	t    HeapType
}
