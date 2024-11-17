package utils

import (
	"fmt"
	"strings"
)

func Duration(t float64) string {
	time := int(t)

	if time < 60 {
		if time == 1 {
			return "1 second"
		}

		return fmt.Sprintf("%d seconds", time)
	}

	if time == 60 {
		return "1 minute"
	}

	output := ""

	seconds := time % 60
	if seconds == 1 {
		output = fmt.Sprintf("%d second %s", seconds, output)
	} else if seconds > 1 {
		output = fmt.Sprintf("%d seconds %s", seconds, output)
	}

	minutes := (time / 60) % 60
	if minutes == 1 {
		output = fmt.Sprintf("%d minute %s", minutes, output)
	} else if minutes > 1 {
		output = fmt.Sprintf("%d minutes %s", minutes, output)
	}

	hours := (time / 3600) % 24
	if hours == 1 {
		output = fmt.Sprintf("%d hour %s", hours, output)
	} else if hours > 1 {
		output = fmt.Sprintf("%d hours %s", hours, output)
	}

	days := (time / 86400) % 86400
	if days == 1 {
		output = fmt.Sprintf("%d day %s", days, output)
	} else if days > 1 {
		output = fmt.Sprintf("%d days %s", days, output)
	}

	return strings.Trim(output, " ")
}
