package queue

type QueueNode struct {
	Value interface{}
	Next  *QueueNode
}

// Queue generic stack implementation using a Singly Linked List
type Queue struct {
	Head   *QueueNode
	Tail   *QueueNode
	Length int
}

// New Create a new queue
// O(1)
func New() *Queue {
	return &Queue{
		Head:   nil,
		Tail:   nil,
		Length: 0,
	}
}

// Dequeue Take the next item off the front of the queue
// O(1)
func (q *Queue) Dequeue() interface{} {
	if q.Head == nil {
		return nil
	}

	front := q.Peek()
	if q.Head == q.Tail {
		q.Head = nil
		q.Tail = nil
	} else {
		q.Head = q.Head.Next
	}

	q.Length--
	return front
}

// Enqueue Put an item on the end of a queue
// O(1)
func (q *Queue) Enqueue(value interface{}) {
	t := &QueueNode{
		Next:  nil,
		Value: value,
	}

	if q.Head == nil {
		q.Head = t
		q.Tail = t
	} else {
		q.Tail.Next = t
		q.Tail = t
	}

	q.Length++
}

// Len Return the number of items in the queue
// O(1)
func (q *Queue) Len() int {
	return q.Length
}

// Peek Return the first item in the queue without removing it
// O(1)
func (q *Queue) Peek() interface{} {
	if q.Head == nil {
		return nil
	}

	return q.Head.Value
}
