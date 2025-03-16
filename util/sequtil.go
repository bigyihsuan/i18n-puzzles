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

func Map[T, U any](s iter.Seq[T], f func(T) U) iter.Seq[U] {
	return func(yield func(U) bool) {
		next, stop := iter.Pull(s)
		defer stop()
		for {
			t, ok := next()
			if !ok {
				return
			}
			u := f(t)
			if !yield(u) {
				return
			}
		}
	}
}
