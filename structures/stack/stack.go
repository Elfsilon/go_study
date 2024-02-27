package stack

import "errors"

type Stack[T any] struct {
	top    *entry[T]
	length int
}

type entry[T any] struct {
	value T
	prev  *entry[T]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		top:    nil,
		length: 0,
	}
}

func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("cannot Peek(), stack is empty")
	}

	return s.top.value, nil
}

func (s *Stack[T]) Pop() error {
	if s.IsEmpty() {
		return errors.New("cannot Pop(), stack is empty")
	}

	s.top = s.top.prev
	s.length--
	return nil
}

func (s *Stack[T]) Push(e T) {
	newEntry := &entry[T]{
		value: e,
		prev:  s.top,
	}
	s.top = newEntry
	s.length++
}

func (s *Stack[T]) PushAll(e ...T) {
	for _, curr := range e {
		s.Push(curr)
	}
}

func (s *Stack[T]) Len() int {
	return s.length
}

func (s *Stack[T]) IsEmpty() bool {
	return s.length == 0
}
