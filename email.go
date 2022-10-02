// Copyright 2022 Emmanuel Ay. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package anonymize

import "strings"

// Email expects an e-mail address and returns an anonymized string.
func Email(input string) string {
	return processEmail(input, '*')
}

// EmailWithCustomRune anonymizes the input string. The 'anonRune' parameter is to use your custom rune to hide characters.
func EmailWithCustomRune(input string, anonRune rune) string {
	return processEmail(input, anonRune)
}

func processEmail(input string, anonRune rune) string {
	if len(input) == 0 || len(strings.TrimSpace(input)) == 0 {
		return ""
	}

	atSplit := strings.Split(input, "@")

	for atIdx, atPart := range atSplit {
		if len(atSplit) > 1 && atIdx == len(atSplit)-1 {
			atSplit[atIdx] = processDomain(atPart, anonRune)
		} else {
			dotSplit := strings.Split(atPart, ".")
			for dotIdx, dotPart := range dotSplit {
				dotSplit[dotIdx] = processWord(dotPart, false, anonRune)
			}
			atSplit[atIdx] = strings.Join(dotSplit, ".")
		}
	}

	return strings.Join(atSplit, "@")
}
