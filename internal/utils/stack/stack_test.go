package stack_test

import (
	"testing"

	"github.com/IAmRadek/rules/internal/utils/stack"
)

func TestStack(t *testing.T) {
	// Create a new stack
	s := stack.Stack[int]{}

	// Test IsEmpty() on an empty stack
	if !s.IsEmpty() {
		t.Errorf("Expected stack to be empty, but it is not.")
	}

	// Test Push() and Peek() on an empty stack
	s.Push(1)
	if value, ok := s.Peek(); !ok || value != 1 {
		t.Errorf("Expected Peek() to return 1, but got %v", value)
	}

	// Test Push() and Pop() on a non-empty stack
	s.Push(2)
	value, ok := s.Pop()
	if !ok || value != 2 {
		t.Errorf("Expected Pop() to return 2, but got %v", value)
	}

	// Test Pop() on an empty stack
	s.Pop()
	value, ok = s.Pop()
	if ok || value != 0 {
		t.Errorf("Expected Pop() on empty stack to return false, but got true")
	}

	// Test MustPop() on an empty stack
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected MustPop() on empty stack to panic, but it did not")
		}
	}()
	s.MustPop()
}
