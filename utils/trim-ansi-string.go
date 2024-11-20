package utils

import (
	"strconv"
)

func TrimAnsiString(str string, limit int) string {
	output := ""
	charsToAdd := limit

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
			escapeLevel -= 1
			output += string(char)

			isEscapeClosed = true
			continue
		}

		output += string(char)

		if escapeLevel == 0 {
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
