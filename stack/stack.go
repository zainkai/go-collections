package stack

// Stack generic stack implementation using a array
type Stack struct {
}

// New Create a new Stack
// O(1)
func New() *Stack {
	return &Stack{}
}

// Len Return the number of items in the Stack
// O(1)
func (s *Stack) Len() int {
	return 0
}

// Peek View the top item on the Stack
// O(1)
func (s *Stack) Peek() interface{} {
	return nil
}

// Pop the top item of the Stack and return it
// O(1)
func (s *Stack) Pop() interface{} {
	return nil
}

// Push a value onto the top of the Stack
// O(1)
func (s *Stack) Push(value interface{}) {

}
