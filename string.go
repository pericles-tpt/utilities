package utilities

import (
	"fmt"
	"strings"
	"unicode"
)

// RemoveCharsIn removes any chars from `cutset` in `target`
func RemoveCharsIn(target string, cutset string) string {
	var (
		ret       = []rune(target)
		cutsetArr = []rune(cutset)
		addIdx    = 0
	)
	for _, ch := range target {
		if IndexOf(ch, cutsetArr) == -1 {
			ret[addIdx] = ch
			addIdx++
		}
	}
	numRemoved := (len(target) - addIdx)
	return string(ret[:len(target)-numRemoved])
}

// AddCharAround adds `char` before and after `target`
func AddCharAround(target string, char rune) string {
	return fmt.Sprintf(`%c%s%c`, char, target, char)
}

// TrimCharsAround trims any chars found in `cutset` from `target`
//
// if `trimUntilNoChange` is true, continues trimming until the string stops changing
func TrimCharsAround(target string, cutset string, trimUntilNoChange bool) string {
	var (
		isNoChange bool
		curr       = target
	)
	for !isNoChange {
		beforeTrim := curr
		curr = strings.Trim(curr, cutset)

		isNoChange = true
		if trimUntilNoChange {
			isNoChange = beforeTrim == curr
		}
	}

	return curr
}

// TrimWhitespaceAround trims any whitespace around `target`
//
// if `trimUntilNoChange` is true, continues trimming until the string stops changing
func TrimWhitespaceAround(target string, trimUntilNoChange bool) string {
	var (
		isNoChange bool
		curr       = target
	)
	for !isNoChange {
		beforeTrim := curr
		curr = strings.TrimSpace(curr)

		isNoChange = true
		if trimUntilNoChange {
			isNoChange = beforeTrim == curr
		}
	}

	return curr
}

// Removes all whitespace around a list of strings
func TrimWhitespaceAroundStrings(targets []string) []string {
	ret := make([]string, len(targets))
	for i, s := range targets {
		ret[i] = TrimWhitespaceAround(s, true)
	}
	return ret
}

// Removes any empty strings found in `targets`
func RemoveEmptyStrings(targets []string) []string {
	var (
		ret    = make([]string, len(targets))
		addIdx = 0
	)
	for i := 0; i < len(targets); i++ {
		ret[addIdx] = targets[i]
		if len(ret[addIdx]) != 0 {
			addIdx++
		}
	}
	emptyCount := len(targets) - addIdx
	return ret[:len(targets)-emptyCount]
}

func GetUpperChars(s string) string {
	var ret string
	for _, c := range s {
		if unicode.IsUpper(c) {
			ret += string(c)
		}
	}
	return ret
}

func GetLowerChars(s string) string {
	var ret string
	for _, c := range s {
		if unicode.IsLower(c) {
			ret += string(c)
		}
	}
	return ret
}
