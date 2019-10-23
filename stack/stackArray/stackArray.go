package stackarray

// StackArray generic stack implementation using a array
type StackArray struct {
	Data []interface{}
}

// New Create a new StackArray
// O(1)
func New() *StackArray {
	return &StackArray{
		make([]interface{}, 0),
	}
}

// Len Return the number of items in the StackArray
// O(1)
func (s *StackArray) Len() int {
	return len(s.Data)
}

// Peek View the top item on the StackArray
// O(1)
func (s *StackArray) Peek() interface{} {
	length := s.Len()
	if length == 0 {
		return nil
	}

	return s.Data[length-1]
}

// Pop the top item of the StackArray and return it
// O(N), this creates a new slice of data every time its called
func (s *StackArray) Pop() interface{} {
	top := s.Peek()
	if top == nil {
		return nil
	}

	length := s.Len()
	s.Data = s.Data[:length-1]

	return top
}

// Push a value onto the top of the StackArray
// O(1) Best case, O(N) worse case
func (s *StackArray) Push(value interface{}) {
	s.Data = append(s.Data, value)
}
