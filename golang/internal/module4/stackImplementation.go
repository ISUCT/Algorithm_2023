package module4

import (
	"errors"
)

type Text interface {
	rune | string
}

type Stack[T Text] struct {
	top      int
	elements []T
}

func (s *Stack[T]) Push(element T) {
	s.top += 1
	s.elements[s.top-1] = element
}

func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		return *new(T), errors.New("Stack underflow")
	} else {
		s.top -= 1
		return s.elements[s.top], nil
	}
}

func (s *Stack[T]) IsEmpty() bool {
	return s.top == 0
}

func (s *Stack[T]) Top() T {
	if !s.IsEmpty() {
		return s.elements[s.top-1]
	}

	return *new(T)
}

func (s *Stack[T]) Len() int {
	return s.top
}

func InitStack[T Text](size int) Stack[T] {
	var emptyStack Stack[T]
	emptyStack.elements = make([]T, size)
	return emptyStack
}
