package utils

import (
	"strconv"
)

func TrimAnsiString(str string, limit int) string {
	output := ""
	charsToAdd := limit

	if charsToAdd <= 0 {
		return output
	}

	escapeLevel := 0
	isEscapeClosed := true

	for index, char := range str {
		if char == '\x1b' {
			escapeLevel += 1
			output += string(char)
			isEscapeClosed = false
			continue
		}

		if escapeLevel > 0 && !isEscapeClosed && char == 'm' {
			output += string(char)

			if str[index-1] == 48 && str[index-2] == 91 {
				// Since this ANSI escape has increased the level as well,
				// the level should be decreased by two here.
				escapeLevel -= 2
			}

			isEscapeClosed = true
			continue
		}

		if charsToAdd > 0 || !isEscapeClosed {
			output += string(char)
		}

		if isEscapeClosed {
			charsToAdd -= 1

			if escapeLevel == 0 && charsToAdd <= 0 {
				break
			}
		}
	}

	output, err := strconv.Unquote(`"` + output + `"`)

	if err != nil {
		return str
	}

	return output
}
