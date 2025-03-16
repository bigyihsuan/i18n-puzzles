package day09

import (
	"bigyihsuan/i18n-puzzles/util"
	"fmt"
	"iter"
	"maps"
	"slices"
	"strings"
	"time"
)

func Solution(input string) string {
	entries := parseInput(input)
	fmt.Println(entries)
	names := make(map[string]struct{})

	for name, dates := range entries {
		for _, layout := range layouts {
			if usesLayout(dates, layout) {
				times := util.Map(slices.Values(dates), func(s string) time.Time {
					t, err := time.Parse(layout, s)
					if err != nil {
						panic(err)
					}
					return t
				})
				if has911(times) {
					names[name] = struct{}{}
				}
			}
		}
	}

	return strings.Join(slices.Sorted(maps.Keys(names)), " ")
}

var layouts = []string{
	"01-02-06",
	"02-01-06",
	"06-01-02",
	"06-02-01",
}

func has911(dates iter.Seq[time.Time]) bool {
	return util.Any(dates, func(date time.Time) bool {
		return date.Year() == 2001 && date.Month() == time.September && date.Day() == 11
	})
}

func usesLayout(dates []string, layout string) bool {
	return util.All(slices.Values(dates), func(date string) bool {
		_, err := time.Parse(layout, date)
		return err == nil
	})
}

type entries map[string][]string

func parseInput(input string) entries {
	entries := make(map[string][]string)
	lines := util.ToSeqNoNewline[string](input)
	for line := range lines {
		split := strings.Split(line, ": ")
		date, names := split[0], strings.Split(split[1], ", ")
		for _, name := range names {
			if _, ok := entries[name]; !ok {
				entries[name] = []string{date}
			} else {
				entries[name] = append(entries[name], date)
			}
		}
	}
	return entries
}
