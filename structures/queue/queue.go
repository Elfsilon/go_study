package queue

import (
	"errors"
	"fmt"
	"strings"
)

type Queue[T any] struct {
	first  *entry[T]
	last   *entry[T]
	length int
}

type entry[T any] struct {
	value T
	prev  *entry[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		first:  nil,
		last:   nil,
		length: 0,
	}
}

func (q *Queue[T]) Enqueue(e T) {
	newEntry := &entry[T]{
		value: e,
		prev:  nil,
	}

	if q.IsEmpty() {
		q.first = newEntry
		q.last = newEntry
	} else {
		q.last.prev = newEntry
		q.last = q.last.prev
	}

	q.length++
}

func (q *Queue[T]) EnqueueAll(values ...T) {
	for _, value := range values {
		q.Enqueue(value)
	}
}

func (q *Queue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, errors.New("cannot Dequeue(), queue is empty")
	}

	res := q.first.value

	if q.length == 1 {
		q.last = nil
	}
	q.first = q.first.prev
	q.length--

	return res, nil
}

func (q *Queue[T]) IsEmpty() bool {
	return q.length == 0
}

func (q *Queue[T]) String() string {
	if q.IsEmpty() {
		return "Que{}"
	}

	var b strings.Builder
	b.WriteString("Que{")
	top := q.first

	counter := q.length - 1
	for top != nil {
		b.WriteString(fmt.Sprint(top.value))
		if counter > 0 {
			b.WriteString("<-")
		}
		top = top.prev
		counter--
	}

	b.WriteString("}")

	return b.String()
}
