package anonymize

import "unicode"

func processWord(input string, allowSingleToken bool, anontoken rune) string {

	if len(input) == 0 {
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
