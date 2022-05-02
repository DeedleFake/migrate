package util

import "golang.org/x/exp/constraints"

// OrderedKey is a type that can both be used as a map key and has a
// natural ordering.
type OrderedKey interface {
	comparable
	constraints.Ordered
}

// SortedKeys returns the keys of m in their naturally sorted order.
func SortedKeys[K OrderedKey, V any, M ~map[K]V](m M) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = SortedInsert(keys, k)
	}
	return keys
}
