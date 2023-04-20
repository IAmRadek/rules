package stack

type Stack[T any] struct {
	data []T
	size int
}

// IsEmpty check if stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return s.size == 0
}

// Push a new value onto the stack
func (s *Stack[T]) Push(str T) {
	if s.size == len(s.data) {
		// If the underlying slice is full, double its capacity.
		newCap := 2 * s.size
		if newCap == 0 {
			newCap = 8 // Start with a small capacity.
		}
		newData := make([]T, newCap)
		copy(newData, s.data)
		s.data = newData
	}
	s.data[s.size] = str
	s.size++
}

// Peek returns the top element of the stack without removing it.
func (s *Stack[T]) Peek() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}
	return s.data[s.size-1], true
}

// Pop and return top element of stack. Return false if stack is empty.
func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}
	s.size--
	return s.data[s.size], true
}

func (s *Stack[T]) MustPop() T {
	v, ok := s.Pop()
	if !ok {
		panic("stack is empty")
	}
	return v
}
