package util

import "iter"

func All[T any](s iter.Seq[T], check func(T) bool) bool {
	for e := range s {
		if !check(e) {
			return false
		}
	}
	return true
}

func Any[T any](s iter.Seq[T], check func(T) bool) bool {
	for e := range s {
		if check(e) {
			return true
		}
	}
	return false
}
