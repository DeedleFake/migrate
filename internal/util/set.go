package util

import "golang.org/x/exp/maps"

// Set is an unordered collection of unique objects. A zero value of
// set is ready to use. Sets are not thread-safe.
type Set[T comparable] struct {
	m map[T]struct{}
}

func (s *Set[T]) init() {
	if s.m == nil {
		s.m = make(map[T]struct{})
	}
}

// Len returns the number of elements in the set.
func (s Set[T]) Len() int {
	return len(s.m)
}

// Add adds an element to the set. It returns true if the element was
// in fact a new element that didn't previously exist in the set.
func (s *Set[T]) Add(v T) bool {
	if s.Contains(v) {
		return false
	}

	s.init()
	s.m[v] = struct{}{}
	return true
}

// AddSet adds all of the elements of s2 to s. It returns true if any
// new elements were added.
func (s *Set[T]) AddSet(s2 *Set[T]) bool {
	var r bool
	for v := range s2.m {
		r = s.Add(v) || r
	}
	return r
}

// Remove removes the element v from s.
func (s *Set[T]) Remove(v T) {
	s.init()
	delete(s.m, v)
}

// Remove removes all of the elements in s2 from s.
func (s *Set[T]) RemoveSet(s2 *Set[T]) {
	s.init()
	for v := range s2.m {
		delete(s.m, v)
	}
}

// Contains returns true if s contains v.
func (s Set[T]) Contains(v T) bool {
	_, ok := s.m[v]
	return ok
}

// Do calls f for each element in s in an undefined order.
func (s Set[T]) Do(f func(T)) {
	for v := range s.m {
		f(v)
	}
}

// Slice returns the elements of s as a slice. The order of the
// contents of the slice is undefined.
func (s Set[T]) Slice() []T {
	return maps.Keys(s.m)
}
