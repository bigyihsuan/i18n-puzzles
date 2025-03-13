package day03

import (
	"bigyihsuan/i18n-puzzles/util"
	_ "embed"
	"errors"
	"fmt"
	"slices"
	"strings"
	"unicode"
)

func Solution(input string) (valid int) {
	lines := util.ToSeq[string](input)

	for password := range lines {
		ok, reasons := isValidPassword(password)
		fmt.Printf("%q => %t\n", password, ok)
		rs := []string{}
		for _, reason := range reasons {
			rs = append(rs, reason.Error())
		}
		if len(rs) != 0 {
			fmt.Println(strings.Join(rs, "; "))
			fmt.Println()
		}
		if ok {
			valid++
		}
	}

	return
}

func isValidPassword(password string) (ok bool, reasons []error) {
	// requirements:
	// - a length of at least 4 and at most 12
	// - at least one digit
	// - at least one uppercase letter (with or without accents, examples: A or Ż)
	// - at least one lowercase letter (with or without accents, examples: a or ŷ)
	// - at least one character that is outside the standard 7-bit ASCII character set (examples: Ű, ù or ř)

	// assume length check is based on chars, not bytes
	withinLen := len([]rune(password)) >= 4 && len([]rune(password)) <= 12
	if !withinLen {
		reasons = append(reasons, fmt.Errorf("not in length range 4 <= n <= 12: %d", len([]rune(password))))
	}

	hasDigit := util.Any(slices.Values([]rune(password)), unicode.IsDigit)
	if !hasDigit {
		reasons = append(reasons, errors.New("missing digit"))
	}

	hasUpperCase := util.Any(slices.Values([]rune(password)), unicode.IsUpper)
	if !hasUpperCase {
		reasons = append(reasons, errors.New("missing uppercase"))
	}

	hasLowerCase := util.Any(slices.Values([]rune(password)), unicode.IsLower)
	if !hasLowerCase {
		reasons = append(reasons, errors.New("missing lowercase"))
	}

	hasNonASCII := util.Any(slices.Values([]rune(password)), func(r rune) bool { return r > unicode.MaxASCII })
	if !hasNonASCII {
		reasons = append(reasons, errors.New("missing non-ascii"))
	}

	return withinLen && hasDigit && hasUpperCase && hasLowerCase && hasNonASCII, reasons
}

//go:embed test-input.txt
var TestInput string

//go:embed input.txt
var Input string
