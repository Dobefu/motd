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

	for _, char := range str {
		if char == '\x1b' {
			escapeLevel += 1
			output += string(char)
			isEscapeClosed = false
			continue
		}

		if escapeLevel > 0 && !isEscapeClosed && char == 'm' {
			output += string(char)

			if isEscapeClosed {
				escapeLevel -= 1
			}

			isEscapeClosed = true
			continue
		}

		output += string(char)

		if escapeLevel == 0 && isEscapeClosed {
			charsToAdd -= 1

			if charsToAdd <= 0 {
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
