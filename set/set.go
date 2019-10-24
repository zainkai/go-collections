package set

// Empty 0 byte size struct
type Empty struct{}

// Set generic hashmap implementation
type Set struct {
	Data map[interface{}]Empty
}

// New create Set
func New() *Set {
	return &Set{
		Data: make(map[interface{}]Empty),
	}
}

// Add insert key to set
// O(1)
func (s *Set) Add(key interface{}) {
	s.Data[key] = Empty{}
}

// Has check if key is in set
// O(1)
func (s *Set) Has(key interface{}) bool {
	_, ok := s.Data[key]
	return ok
}

// Delete remove key from set
// O(1)
func (s *Set) Delete(key interface{}) {
	delete(s.Data, key)
}

// Len return number of elements in set
// O(1)
func (s *Set) Len() int {
	return len(s.Data)
}

// Keys generate array of keys in set, unordered
// O(N)
func (s *Set) Keys() []interface{} {
	result := make([]interface{}, 0, 0)

	for key := range s.Data {
		result = append(result, key)
	}

	return result
}

// ForEach iterate over all elements in set, unordered
// callback func(key interface{})
// O(N)
func (s *Set) ForEach(callback func(key interface{})) {
	for key := range s.Data {
		callback(key)
	}
}
