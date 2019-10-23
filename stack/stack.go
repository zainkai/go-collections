package stack

type Stack struct {
}

// Create a new stack
func New() *Stack {
	return &Stack{}
}

// Return the number of items in the stack
func (this *Stack) Len() int {
	return 0
}

// View the top item on the stack
func (this *Stack) Peek() interface{} {
	return nil
}

// Pop the top item of the stack and return it
func (this *Stack) Pop() interface{} {
	return nil
}

// Push a value onto the top of the stack
func (this *Stack) Push(value interface{}) {

}
