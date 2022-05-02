package util

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

// SortedInsert inserts e into the sorted slice s such that the newly
// returned slice containing e will maintain the natual sort order of
// E.
func SortedInsert[E constraints.Ordered, S ~[]E](s S, e E) S {
	i, _ := slices.BinarySearch(s, e)
	return slices.Insert(s, i, e)
}

// SortedInsertFunc is like SortedInsert but uses the given comparison
// function to determine the sorting order.
func SortedInsertFunc[E any, S ~[]E](s S, e E, cmp func(E, E) int) S {
	i, _ := slices.BinarySearchFunc(s, e, cmp)
	return slices.Insert(s, i, e)
}
