package queue

type Queue struct {
}

// Create a new queue
func New() *Queue {
	return &Queue{}
}

// Take the next item off the front of the queue
func (this *Queue) Dequeue() interface{} {
	return nil
}

// Put an item on the end of a queue
func (this *Queue) Enqueue(value interface{}) {
}

// Return the number of items in the queue
func (this *Queue) Len() int {
	return 0
}

// Return the first item in the queue without removing it
func (this *Queue) Peek() interface{} {
	return nil
}
