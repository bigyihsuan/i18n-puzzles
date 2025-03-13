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

type trip struct {
	Start, End zoneTime
}

type zoneTime struct {
	Zone, Time string
}

func (z zoneTime) AsTime() (time.Time, error) {
	l, err := time.LoadLocation(z.Zone)
	if err != nil {
		return time.Now(), err
	}
	return time.ParseInLocation("Jan 02, 2006, 15:04", z.Time, l)
}

type locKind string

const (
	lcDeparture locKind = "Departure:"
	lcArrival   locKind = "Arrival:"
)

func newLoc(input string, kind locKind) zoneTime {
	var l zoneTime
	s, _ := strings.CutPrefix(input, string(kind))
	fields := strings.Fields(s)
	l.Zone = fields[0]
	l.Time = strings.Join(fields[1:], " ")
	return l
}

func parseInput(input string) (ts []trip) {
	// input is grouped in sets of 3 lines
	lines := slices.Collect(util.ToSeq[string](input))
	tripLines := slices.Chunk(lines, 3)
	for tl := range tripLines {
		departure := newLoc(tl[0], lcDeparture)
		arrival := newLoc(tl[1], lcArrival)
		ts = append(ts, trip{departure, arrival})
	}

	return
}

//go:embed test-input.txt
var TestInput string

//go:embed input.txt
var Input string
