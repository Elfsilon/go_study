package search

import "cmp"

type SearchFn[T cmp.Ordered] func(target T, where []T) (int, error)

func Linear[T cmp.Ordered](target T, where []T) (int, error) {
	for i := 0; i < len(where); i++ {
		if where[i] == target {
			return i, nil
		}
	}

	return -1, nil
}

func Binary[T cmp.Ordered](target T, where []T) (int, error) {
	l := len(where)
	if l == 0 {
		return -1, nil
	}

	m := l / 2

	if where[m] == target {
		return m, nil
	}

	if target > where[m] {
		return Binary(target, where[m+1:])
	} else {
		return Binary(target, where[0:m])
	}
}
