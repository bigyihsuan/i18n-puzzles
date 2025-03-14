package day07

import (
	"bigyihsuan/i18n-puzzles/util"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func Solution(input string) (out int) {
	audits := parseInput(input)
	for i, audit := range audits {
		zoneName, offset := audit.Timestamp.Zone()
		fixed := audit.Fixed()
		fmt.Printf("%s %q %d\n", audit, zoneName, offset/60/60)
		fmt.Println(fixed.Format(util.TimeFormat))
		// fmt.Println(Expected[i])
		fmt.Println()
		out += fixed.Hour() * (i + 1)
	}
	return out
}

const (
	halifax  = "America/Halifax"
	santiago = "America/Santiago"
)

type audit struct {
	Timestamp          time.Time
	Correct, Incorrect int
}

func (a audit) Fixed() time.Time {
	halLoc, _ := time.LoadLocation(halifax)
	santLoc, _ := time.LoadLocation(santiago)

	hal := a.Timestamp.In(halLoc)
	sant := a.Timestamp.In(santLoc)

	if hal.Hour() == a.Timestamp.Hour() {
		return hal.Add(-time.Duration(a.Incorrect) * time.Minute).Add(time.Duration(a.Correct) * time.Minute)
	} else {
		return sant.Add(-time.Duration(a.Incorrect) * time.Minute).Add(time.Duration(a.Correct) * time.Minute)
	}
}

func (a audit) String() string {
	return fmt.Sprintf("{%s, %d, %d}", a.Timestamp, a.Correct, a.Incorrect)
}

func parseInput(input string) (entries []audit) {
	lines := util.ToSeqNoNewline[string](input)
	for line := range lines {
		fields := strings.Fields(line)
		timestamp, err := time.ParseInLocation(util.TimeFormat, fields[0], time.UTC)
		if err != nil {
			panic(err)
		}
		correct, err := strconv.Atoi(fields[1])
		if err != nil {
			panic(err)
		}
		incorrect, err := strconv.Atoi(fields[2])
		if err != nil {
			panic(err)
		}
		entries = append(entries, audit{
			Timestamp: timestamp,
			Correct:   correct,
			Incorrect: incorrect,
		})
	}
	return
}
