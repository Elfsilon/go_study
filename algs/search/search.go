package search

import "cmp"

type SearchFn[T cmp.Ordered] func(target T, where []T) int

func Linear[T cmp.Ordered](target T, where []T) int {
	for i := 0; i < len(where); i++ {
		if where[i] == target {
			return i
		}
	}

	return -1
}

func Binary[T cmp.Ordered](target T, where []T) int {
	l := len(where)
	if l == 0 {
		return -1
	}

	m := l / 2

	if where[m] == target {
		return m
	}

	if target > where[m] {
		return Binary(target, where[m+1:])
	} else {
		return Binary(target, where[0:m])
	}
}

func BinaryIter[T cmp.Ordered](target T, where []T) int {
	if len(where) == 0 {
		return -1
	}

	l := 0
	h := len(where) - 1

	for l <= h {
		m := (l + h) / 2
		if where[m] == target {
			return m
		}

		if target > where[m] {
			l = m + 1
		} else {
			h = m - 1
		}
	}

	return -1
}