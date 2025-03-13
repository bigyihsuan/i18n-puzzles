package day02

import (
	"bigyihsuan/i18n-puzzles/util"
	_ "embed"
	"fmt"
	"iter"
	"maps"
	"time"
)

const timeFormat = "2006-01-02T15:04:05-07:00" // RFC3339 with `-` instead of `Z`

func Solution(input string) string {
	waves := make(map[time.Time]int)

	lines := util.ToSeq[string](input)
	times := toTimes(lines)

	// find occurrences
	for t := range times {
		// fmt.Println(t.Format(TIME_FORMAT))
		t = t.In(time.UTC) // normalize to +00:00
		if _, ok := waves[t]; !ok {
			waves[t] = 1
		} else {
			waves[t]++
		}
	}

	// pick the first wave that happened 4+ times
	for t, n := range maps.All(waves) {
		if n >= 4 {
			return t.Format(timeFormat)
		}
	}

	return time.Now().Format(timeFormat)
}

//go:embed test-input.txt
var TestInput string

//go:embed input.txt
var Input string

func toTimes[T ~string](lines iter.Seq[T]) iter.Seq[time.Time] {
	return func(yield func(time.Time) bool) {
		next, stop := iter.Pull(lines)
		defer stop()
		for {
			line, ok := next()
			if !ok {
				return
			}
			t, err := time.Parse(timeFormat, string(line))
			if err != nil {
				panic(fmt.Errorf("LinesToTimes time parsing failed: %w", err))
			}
			if !yield(t) {
				return
			}
		}
	}
}
