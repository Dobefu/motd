package utils

import "fmt"

func FormatBytes(bytes int) string {
	newBytes := float64(bytes)
	iteration := 0

	units := []string{
		"bytes",
		"KiB",
		"MiB",
		"GiB",
		"TiB",
		"PiB",
		"EiB",
	}

	for newBytes > 1024 {
		iteration += 1
		newBytes = newBytes / 1024
	}

	return fmt.Sprintf("%d%s", int(newBytes), units[iteration])
}
