// Package Anonymize is a small and simple package for anonymizing names, e-mails and domains.
//
// Copyright 2022 Emmanuel Ay. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package anonymize

import "strings"

// Domain expects a domain name and anonymizes the input string
func Domain(input string) string {
	return processDomain(input, '*')
}

// DomainWithCustomRune expects a domain name and anonymizes the input string. The 'anonRune' parameter is to use your custom rune to hide characters.
func DomainWithCustomRune(input string, anonRune rune) string {
	return processDomain(input, anonRune)
}

func processDomain(input string, anonRune rune) string {

	dotSplit := strings.Split(input, ".")

	for dotIdx, dotPart := range dotSplit {

		dashSplit := strings.Split(dotPart, "-")
		for dashIdx, dashPart := range dashSplit {
			dashSplit[dashIdx] = processWord(dashPart, false, anonRune)
		}

		if len(dotSplit) > 1 && dotIdx == len(dotSplit)-1 || strings.EqualFold(strings.ToLower(dotPart), "www") {
			dotSplit[dotIdx] = dotPart
		} else {
			dotSplit[dotIdx] = strings.Join(dashSplit, "-")
		}

	}

	return strings.Join(dotSplit, ".")
}
