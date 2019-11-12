package frutils

import "strings"

func RemoveAllCharactersBefore(str string, separator string) string {
	hasAppeared := false
	result := ""

	for _, c := range str {
		if string(c) == separator {
			hasAppeared = true
		}

		if hasAppeared {
			result += string(c)
		}
	}

	return result
}

func RemoveNewLinesAndWhiteSpace(str string) string {
	return strings.Trim(strings.Trim(str, "\n"), `"`)
}

func RemoveExteriorSquareBrackets(str string) string {
	return strings.TrimSuffix(strings.TrimSuffix(strings.TrimPrefix(str, "["), "\n"), "]")
}
