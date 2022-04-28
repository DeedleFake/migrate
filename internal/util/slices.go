package util

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

func SortedInsert[E constraints.Ordered, S ~[]E](s S, e E) S {
	i, _ := slices.BinarySearch(s, e)
	return slices.Insert(s, i, e)
}

func SortedInsertFunc[E any, S ~[]E](s S, e E, cmp func(E, E) int) S {
	i, _ := slices.BinarySearchFunc(s, e, cmp)
	return slices.Insert(s, i, e)
}
