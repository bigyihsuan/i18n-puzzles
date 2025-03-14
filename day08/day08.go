package day08

import (
	"bigyihsuan/i18n-puzzles/util"
	"errors"
	"fmt"
	"regexp"
	"slices"
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
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
	// - at least one accented or unaccented vowel (a, e, i, o, u) (examples: i, Á or ë).
	// - at least one accented or unaccented consonant, examples: s, ñ or ŷ
	// - no recurring letters in any form. Ignoring accents and case, letters should not recur.
	// 		For example, in 'niña' the 'n' occurs twice, one time with accent and one time without.
	// 		'Usul' is out because the 'u' occurs twice, first uppercase and then lowercase.

	ok = true

	// assume length check is based on chars, not bytes
	withinLen := len([]rune(password)) >= 4 && len([]rune(password)) <= 12
	if !withinLen {
		reasons = append(reasons, fmt.Errorf("not in length range 4 <= n <= 12: %d", len([]rune(password))))
	}
	ok = ok && withinLen

	hasDigit := util.Any(slices.Values([]rune(password)), unicode.IsDigit)
	if !hasDigit {
		reasons = append(reasons, errors.New("missing digit"))
	}
	ok = ok && hasDigit

	hasVowel := util.Any(slices.Values([]rune(password)), func(r rune) bool {
		r = unicode.ToLower(r)
		normalized := []rune(string(norm.NFD.Bytes(runeToBytes(r))))
		// fmt.Printf("%q\n", normalized)
		return regexp.MustCompile(`[aeiou]`).MatchString(string(normalized[0]))
		// first normalized rune is the base char
	})
	if !hasVowel {
		reasons = append(reasons, errors.New("missing vowel"))
	}
	ok = ok && hasVowel

	hasConsonant := util.Any(slices.Values([]rune(password)), func(r rune) bool {
		r = unicode.ToLower(r)
		normalized := []rune(string(norm.NFD.Bytes(runeToBytes(r))))
		return regexp.MustCompile(`[bcdfghjklmnpqrstvwxyz]`).MatchString(string(normalized[0]))
	})
	if !hasConsonant {
		reasons = append(reasons, errors.New("missing consonant"))
	}
	ok = ok && hasConsonant

	// no dups
	// normalize to lowercase, unaccented
	hasUnique := true
	var sb strings.Builder
	for r := range slices.Values([]rune(strings.ToLower(password))) {
		normalizedRune := []rune(string(norm.NFD.Bytes(runeToBytes(r))))
		sb.WriteRune(normalizedRune[0])
	}
	normalized := sb.String()

	seenRunes := make(map[rune]struct{})
	for r := range slices.Values([]rune(normalized)) {
		if _, ok := seenRunes[r]; !ok {
			seenRunes[r] = struct{}{}
		} else {
			reasons = append(reasons, fmt.Errorf("duplicate character: %q", r))
			hasUnique = false
		}
	}
	ok = ok && hasUnique

	return ok, reasons
}

func runeToBytes(r rune) []byte { return []byte(string(r)) }
