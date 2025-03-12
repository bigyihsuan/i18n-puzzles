package util

import (
	"iter"
	"slices"
	"strings"
)

func ToSeq[T ~string](input string) iter.Seq[T] {
	ls := []T{}
	for line := range strings.Lines(input) {
		ls = append(ls, T(strings.TrimSpace(line)))
	}
	return slices.Values(ls)
}
