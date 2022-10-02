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
	if len(input) == 0 || len(strings.TrimSpace(input)) == 0 {
		return ""
	}

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
