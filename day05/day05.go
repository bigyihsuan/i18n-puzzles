package day05

import (
	"bigyihsuan/i18n-puzzles/util"
	_ "embed"
	"iter"
	"strings"
)

const poo = '💩'

func Solution(input string) (total int) {
	lines := util.ToSeqRaw[string](input)
	next, stop := iter.Pull(lines)
	defer stop()
	row := 0
	col := 0
	for {
		l, ok := next()
		if !ok {
			break
		}
		line := []rune(strings.TrimRight(l, "\n"))
		// fmt.Printf("%d %q\n", row, string(line))
		// fmt.Printf("%q\n", string(line[col]))
		if line[col] == poo {
			total++
		}
		col = (col + 2) % len(line)
		row++
	}
	return total
}

//go:embed test-input.txt
var TestInput string

//go:embed input.txt
var Input string
