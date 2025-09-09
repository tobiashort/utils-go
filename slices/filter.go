package slices

import (
	std_slices "slices"
)

func Filter[T any](s []T, predicate func(T) bool) []T {
	return std_slices.Collect(func(yield func(T) bool) {
		for _, n := range s {
			if predicate(n) {
				if !yield(n) {
					return
				}
			}
		}
	})
}
