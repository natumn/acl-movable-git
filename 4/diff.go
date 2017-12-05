package main

import (
	"fmt"
	"regexp"
	"strings"
)

func diff(a, b string, unified bool) string {
	var ap, bp rune
	if unified {
		ap = '-'
		bp = '+'
	} else {
		ap = '<'
		bp = '>'
	}

	minusA := fmt.Sprintln(regexp.MustCompile(`(^|\n)`).ReplaceAllString(strings.TrimRight(a, "\n"), fmt.Sprintf("$1%c ", ap)))
	plusB := fmt.Sprintln(regexp.MustCompile(`(^|\n)`).ReplaceAllString(strings.TrimRight(b, "\n"), fmt.Sprintf("$1%c ", bp)))
	return fmt.Sprintf("%s%s", minusA, plusB)
}
