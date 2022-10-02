package anonymize

import (
	"strings"
)

// Name anonymizes the input string based on letter case.
func Name(input string) string {
	return processName(input, '*')
}

// NameWithCustomRune anonymizes the input string based on letter case. The 'anonRune' parameter is to use your custom rune to hide characters.
func NameWithCustomRune(input string, anonRune rune) string {
	return processName(input, anonRune)
}

func processName(input string, anonRune rune) string {
	if len(input) == 0 || len(strings.TrimSpace(input)) == 0 {
		return ""
	}

	parts := strings.Split(input, " ")

	for idx, part := range parts {
		parts[idx] = processWord(part, true, anonRune)
	}

	return strings.Join(parts, " ")
}
