package day01

import (
	"bigyihsuan/i18n-puzzles/util"
	_ "embed"
	"fmt"
)

const (
	// SMS = 160 *bytes* = 11 cents
	// tweet = 140 *chars* = 7 cents
	SMS_LIMIT_BYTES   int = 160
	TWEET_LIMIT_CHARS int = 140
)

func Fees(input string) (total int) {
	lines := util.ToSeq[Line](input)

	for line := range lines {
		fmt.Printf("%q\nbytes: %d\nchars: %d\ncost: %d\n\n", line, len([]byte(line)), len([]rune(line)), line.Cost())
		total += line.Cost()
	}

	return total
}

type Line string

func (l Line) Cost() int {
	// SMS = 11 cents
	// tweet = 7 cents
	// SMS + tweet = 13 cents
	switch {
	case l.IsSMS() && l.IsTweet():
		return 13
	case l.IsSMS():
		return 11
	case l.IsTweet():
		return 7
	default:
		return 0
	}
}

func (l Line) IsSMS() bool {
	return len([]byte(l)) <= SMS_LIMIT_BYTES
}
func (l Line) IsTweet() bool {
	return len([]rune(l)) <= TWEET_LIMIT_CHARS
}

//go:embed testdata.txt
var TestInput string

//go:embed input.txt
var Input string
