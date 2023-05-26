package module4

import (
	"errors"
)

type Number interface {
	int
}

type Text interface {
	rune | string
}

type Element interface {
	Number | Text
}

type Stack[E Element] struct {
	top      int
	elements []E
}

func (s *Stack[E]) Push(element E) {
	s.top += 1
	s.elements[s.top-1] = element
}

func (s *Stack[E]) Pop() (E, error) {
	if s.IsEmpty() {
		return *new(E), errors.New("Stack underflow")
	} else {
		s.top -= 1
		return s.elements[s.top], nil
	}
}

func (s *Stack[E]) IsEmpty() bool {
	return s.top == 0
}

func (s *Stack[E]) Top() E {
	if !s.IsEmpty() {
		return s.elements[s.top-1]
	}

	return *new(E)
}

func (s *Stack[E]) Len() int {
	return s.top
}

func InitStack[E Element](size int) Stack[E] {
	var emptyStack Stack[E]
	emptyStack.elements = make([]E, size)
	return emptyStack
}
