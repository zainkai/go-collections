package stack

type StackNode struct {
	Value interface{}
	Next  *StackNode
}

// Stack generic stack implementation using a Linked List
type Stack struct {
	Length int
	Last   *StackNode
}

// New Create a new Stack
// O(1)
func New() *Stack {
	return &Stack{
		Length: 0,
		Last:   nil,
	}
}

// Len Return the number of items in the Stack
// O(1)
func (s *Stack) Len() int {
	return s.Length
}

// Peek View the top item on the Stack
// O(1)
func (s *Stack) Peek() interface{} {
	return s.Last.Value
}

// Pop the top item of the Stack and return it
// O(1)
func (s *Stack) Pop() interface{} {
	if s.Last == nil {
		return nil
	}

	top := s.Peek()
	s.Last = s.Last.Next

	s.Length--

	return top
}

// Push a value onto the top of the Stack
// O(1)
func (s *Stack) Push(value interface{}) {
	t := &StackNode{
		Value: value,
		Next:  s.Last,
	}

	s.Last = t
	s.Length++
}
