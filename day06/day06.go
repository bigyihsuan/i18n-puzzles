package day06

import (
	"bigyihsuan/i18n-puzzles/util"
	"fmt"
	"slices"
	"strings"
	"unicode"

	"golang.org/x/text/encoding/charmap"
)

func Solution(input string) int {
	badDict, crossword := parseInput(input)
	dictionary := fixDictionary(badDict)
	fmt.Printf("%q\n", dictionary)
	fmt.Printf("%q\n", crossword)

	filledIndexes, lengths, trimmed := crosswordIndexesLengths(crossword)
	totalIndexes := 0
	for i := range filledIndexes {
		index := filledIndexes[i]
		targetRune := []rune(trimmed[i])[index]
		length := lengths[i]

		found := slices.IndexFunc(dictionary, func(e string) bool { return len([]rune(e)) == length && []rune(e)[index] == targetRune })
		if found < 0 {
			panic(fmt.Errorf("couldn't find crossword %q in dictionary %q\nindex: %d, target rune: %q, len: %d", crossword[i], dictionary, index, string(targetRune), length))
		}
		fmt.Printf("%d: %q => %q\n", i, crossword[i], dictionary[found])
		totalIndexes += found + 1
	}

	return totalIndexes
}

func crosswordIndexesLengths(crossword []string) (indexes []int, lengths []int, trimmed []string) {
	for _, word := range crossword {
		trimmedWord := strings.TrimSpace(word)
		fmt.Printf("%q\n", trimmedWord)
		indexes = append(indexes, strings.IndexFunc(trimmedWord, func(r rune) bool { return !(r == '.' || unicode.IsSpace(r)) }))
		lengths = append(lengths, len([]rune(trimmedWord)))
		trimmed = append(trimmed, trimmedWord)
	}
	return
}

// file is utf-8
// every 3rd, 5th = utf-8, loaded by iso-latin-1, then stored utf-8 (iso-latin-1 = ISO-8859-1)
// every 15th = doubly miscocded
func fixDictionary(dictionary []string) (newDictionary []string) {
	encoder := charmap.ISO8859_1.NewEncoder()
	for i, entry := range dictionary {
		index := i + 1
		var d string = entry
		var err error
		if index%3 == 0 || index%5 == 0 {
			d, err = encoder.String(d)
			if err != nil {
				panic(err)
			}
		}
		if index%15 == 0 {
			d, err = encoder.String(d)
			if err != nil {
				panic(err)
			}
		}
		fmt.Printf("%d %q => %q\n", index, entry, d)
		newDictionary = append(newDictionary, d)
	}
	return
}

func parseInput(input string) (dictionary []string, crossword []string) {
	parts := strings.Split(input, "\n\n")
	dictionary = slices.Collect(util.ToSeqNoNewline[string](parts[0]))
	crossword = slices.Collect(util.ToSeqNoNewline[string](parts[1]))
	return
}
