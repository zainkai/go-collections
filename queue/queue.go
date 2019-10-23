package queue

// Queue generic stack implementation using a Linked List
type Queue struct {
}

// New Create a new queue
// O(1)
func New() *Queue {
	return &Queue{}
}

// Dequeue Take the next item off the front of the queue
// O(1)
func (q *Queue) Dequeue() interface{} {
	return nil
}

// Enqueue Put an item on the end of a queue
// O(1)
func (q *Queue) Enqueue(value interface{}) {
}

// Len Return the number of items in the queue
// O(1)
func (q *Queue) Len() int {
	return 0
}

// Peek Return the first item in the queue without removing it
// O(1)
func (q *Queue) Peek() interface{} {
	return nil
}
