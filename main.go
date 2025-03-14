package main

import (
	"bigyihsuan/i18n-puzzles/day01"
	"bigyihsuan/i18n-puzzles/day02"
	"bigyihsuan/i18n-puzzles/day03"
	"bigyihsuan/i18n-puzzles/day04"
	"bigyihsuan/i18n-puzzles/day05"
	"bigyihsuan/i18n-puzzles/day06"
	"bigyihsuan/i18n-puzzles/day07"
	"bigyihsuan/i18n-puzzles/day08"
	"flag"
	"fmt"
)

var day = flag.Int("day", 0, "day number, 1-x")
var useTest = flag.Bool("test", false, "use test input")

func main() {
	flag.Parse()

	switch *day {
	case 1:
		if *useTest {
			fmt.Println(day01.Solution(day01.TestInput))
		} else {
			fmt.Println(day01.Solution(day01.Input))
		}
	case 2:
		if *useTest {
			fmt.Println(day02.Solution(day02.TestInput))
		} else {
			fmt.Println(day02.Solution(day02.Input))
		}
	case 3:
		if *useTest {
			fmt.Println(day03.Solution(day03.TestInput))
		} else {
			fmt.Println(day03.Solution(day03.Input))
		}
	case 4:
		if *useTest {
			fmt.Println(day04.Solution(day04.TestInput))
		} else {
			fmt.Println(day04.Solution(day04.Input))
		}
	case 5:
		if *useTest {
			fmt.Println(day05.Solution(day05.TestInput))
		} else {
			fmt.Println(day05.Solution(day05.Input))
		}
	case 6:
		if *useTest {
			fmt.Println(day06.Solution(day06.TestInput))
		} else {
			fmt.Println(day06.Solution(day06.Input))
		}
	case 7:
		if *useTest {
			fmt.Println(day07.Solution(day07.TestInput))
		} else {
			fmt.Println(day07.Solution(day07.Input))
		}
	case 8:
		if *useTest {
			fmt.Println(day08.Solution(day08.TestInput))
		} else {
			fmt.Println(day08.Solution(day08.Input))
		}
	default:
		flag.Usage()
	}
}
