package slicesx

// Unique returns a new slice with no duplicate elements.
func Unique[T comparable](xs []T) []T {
	res := make([]T, 0, len(xs))
	m := make(map[T]struct{}, len(xs))
	for _, x := range xs {
		if _, ok := m[x]; !ok {
			m[x] = struct{}{}
			res = append(res, x)
		}
	}
	return res
}

// Find returns the first element that matches the predicate p.
func Find[T any](xs []T, p func(T) bool) (T, bool) {
	for _, x := range xs {
		if p(x) {
			return x, true
		}
	}

	return *new(T), false
}

// TakeWhile returns the first subset of slice elements the matches the
// predicate p.
func TakeWhile[T any](xs []T, p func(T) bool) []T {
	res := make([]T, 0, len(xs))
	for _, x := range xs {
		if !p(x) {
			break
		}
		res = append(res, x)
	}
	return res
}

// DropWhile drops the first subset of slice elements the matches the predicate
// p.
func DropWhile[T any](xs []T, p func(T) bool) []T {
	res := make([]T, 0)

	i := 0
	for i < len(xs) && p(xs[i]) {
		i += 1
	}

	if i == len(xs) {
		return nil
	}

	return append(res, xs[i:]...)
}

// GroupBy groups slice elements into map by the predicate function f.
func GroupBy[T any, K comparable](xs []T, f func(T) K) map[K][]T {
	m := make(map[K][]T)
	for _, x := range xs {
		k := f(x)
		v, ok := m[k]
		if !ok {
			m[k] = []T{x}
		} else {
			m[k] = append(v, x)
		}
	}
	return m
}

// Group breaks the slice xs into slice of slices with the fixed size n.
func Group[T any](xs []T, n int) [][]T {
	res := make([][]T, 0, len(xs)/n)
	cur := make([]T, 0, n)

	for _, x := range xs {
		if len(cur) == n {
			res = append(res, cur)
			cur = make([]T, 0, n)
		}
		cur = append(cur, x)
	}

	if len(cur) > 0 {
		res = append(res, cur)
	}

	return res
}

// Partition splits the slice into two by the predicate p.
func Partition[T any](xs []T, p func(T) bool) ([]T, []T) {
	positive := make([]T, 0)
	negative := make([]T, 0)

	for _, x := range xs {
		if p(x) {
			positive = append(positive, x)
		} else {
			negative = append(negative, x)
		}
	}

	return positive, negative
}
