package day01

import (
	"bigyihsuan/i18n-puzzles/util"
	_ "embed"
	"fmt"
)

const (
	// SMS = 160 *bytes* = 11 cents
	// tweet = 140 *chars* = 7 cents
	smsLimitBytes   int = 160
	tweetLimitChars int = 140
)

func Solution(input string) (total int) {
	lines := util.ToSeq[line](input)

	for line := range lines {
		fmt.Printf("%q\nbytes: %d\nchars: %d\ncost: %d\n\n", line, len([]byte(line)), len([]rune(line)), line.cost())
		total += line.cost()
	}

	return total
}

type line string

func (l line) cost() int {
	// SMS = 11 cents
	// tweet = 7 cents
	// SMS + tweet = 13 cents
	switch {
	case l.isSMS() && l.isTweet():
		return 13
	case l.isSMS():
		return 11
	case l.isTweet():
		return 7
	default:
		return 0
	}
}

func (l line) isSMS() bool {
	return len([]byte(l)) <= smsLimitBytes
}
func (l line) isTweet() bool {
	return len([]rune(l)) <= tweetLimitChars
}

//go:embed testdata.txt
var TestInput string

//go:embed input.txt
var Input string
