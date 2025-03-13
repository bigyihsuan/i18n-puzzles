package day04

import (
	"bigyihsuan/i18n-puzzles/util"
	_ "embed"
	"fmt"
	"slices"
	"strings"
	"time"
)

func Solution(input string) (totalMinutes int) {
	trips := parseInput(input)
	for trip := range slices.Values(trips) {
		startTime, err := trip.Start.AsTime()
		if err != nil {
			panic(err)
		}
		endTime, err := trip.End.AsTime()
		if err != nil {
			panic(err)
		}
		minutes := int(endTime.Sub(startTime).Minutes())
		fmt.Printf("%s -- %s => %d\n", startTime, endTime, minutes)
		totalMinutes += minutes
	}

	return totalMinutes
}

type Trip struct {
	Start, End ZoneTime
}

type ZoneTime struct {
	Zone, Time string
}

func (z ZoneTime) AsTime() (time.Time, error) {
	l, err := time.LoadLocation(z.Zone)
	if err != nil {
		return time.Now(), err
	}
	return time.ParseInLocation("Jan 02, 2006, 15:04", z.Time, l)
}

type LocKind string

const (
	DEPARTURE LocKind = "Departure:"
	ARRIVAL   LocKind = "Arrival:"
)

func NewLoc(input string, kind LocKind) ZoneTime {
	var l ZoneTime
	s, _ := strings.CutPrefix(input, string(kind))
	fields := strings.Fields(s)
	l.Zone = fields[0]
	l.Time = strings.Join(fields[1:], " ")
	return l
}

func parseInput(input string) (ts []Trip) {
	// input is grouped in sets of 3 lines
	lines := slices.Collect(util.ToSeq[string](input))
	tripLines := slices.Chunk(lines, 3)
	for tl := range tripLines {
		departure := NewLoc(tl[0], DEPARTURE)
		arrival := NewLoc(tl[1], ARRIVAL)
		ts = append(ts, Trip{departure, arrival})
	}

	return
}

//go:embed test-input.txt
var TestInput string

//go:embed input.txt
var Input string
