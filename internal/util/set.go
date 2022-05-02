package util

import "golang.org/x/exp/maps"

type Set[T comparable] struct {
	m map[T]struct{}
}

func (s *Set[T]) init() {
	if s.m == nil {
		s.m = make(map[T]struct{})
	}
}

func (s Set[T]) Len() int {
	return len(s.m)
}

func (s *Set[T]) Add(v T) bool {
	if s.Contains(v) {
		return false
	}

	s.init()
	s.m[v] = struct{}{}
	return true
}

func (s *Set[T]) AddSet(s2 *Set[T]) bool {
	var r bool
	for v := range s2.m {
		r = s.Add(v) || r
	}
	return r
}

func (s *Set[T]) Remove(v T) {
	s.init()
	delete(s.m, v)
}

func (s *Set[T]) RemoveSet(s2 *Set[T]) {
	s.init()
	for v := range s2.m {
		delete(s.m, v)
	}
}

func (s Set[T]) Contains(v T) bool {
	_, ok := s.m[v]
	return ok
}

func (s Set[T]) Do(f func(T)) {
	for v := range s.m {
		f(v)
	}
}

func (s Set[T]) Slice() []T {
	return maps.Keys(s.m)
}
