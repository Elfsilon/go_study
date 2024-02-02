package structures

type Entry[T comparable] struct {
	value T
	next  *Entry[T]
}

func NewEntry[T comparable](value T) *Entry[T] {
	return &Entry[T]{
		value: value,
	}
}

type LinkedList[T comparable] struct {
	Head  *Entry[T]
	Tail  *Entry[T]
	Count int
}

func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (l *LinkedList[T]) Add(value T) {
	entry := NewEntry(value)
	if l.Head == nil {
		l.Head = entry
		l.Tail = l.Head
	} else {
		l.Tail.next = entry
		l.Tail = l.Tail.next
	}

	l.Count += 1
}

func (l *LinkedList[T]) RemoveFirst(target T) {
	current := l.Head
	removedCount := 0

	for current != nil {
		if current.value == target {
			l.Head = current.next
			current.next = nil
		}
		current = current.next
		removedCount += 1
	}

	l.Count -= removedCount
}

func (l *LinkedList[T]) Contains(target T) bool {
	current := l.Head
	for current != nil {

		if current.value == target {
			return true
		}
		current = current.next
	}

	return false
}
