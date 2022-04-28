package util

import "golang.org/x/exp/constraints"

type OrderedKey interface {
	comparable
	constraints.Ordered
}

func SortedKeys[K OrderedKey, V any, M ~map[K]V](m M) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = SortedInsert(keys, k)
	}
	return keys
}
