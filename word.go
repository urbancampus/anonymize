// Package Anonymize is a small and simple package for anonymizing names, e-mails and domains.
//
// Copyright 2022 Emmanuel Ay. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package anonymize

import (
	"strings"
	"unicode"
)

func processWord(input string, allowSingleToken bool, anontoken rune) string {

	if len(input) == 0 || len(strings.TrimSpace(input)) == 0 {
		return ""
	}

	if !allowSingleToken && len(input) == 1 {
		return string(anontoken)
	}

	runeInput := []rune(input)
	runeOutput := make([]rune, 0, len(runeInput))

	var currentRune rune
	var previousRune rune

	for {
		currentRune = runeInput[0]

		if (previousRune == 0) ||
			(!unicode.IsUpper(previousRune) && unicode.IsUpper(currentRune)) {
			runeOutput = append(runeOutput, currentRune)
		} else {
			runeOutput = append(runeOutput, anontoken)
		}

		if len(runeInput) > 1 {
			runeInput = runeInput[1:]
			previousRune = currentRune
			continue
		}

		break
	}

	return string(runeOutput)
}
