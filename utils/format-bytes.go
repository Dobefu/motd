package utils

import "fmt"

func FormatBytes(bytes int) string {
	new_bytes := float64(bytes)
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

	for new_bytes > 1024 {
		iteration += 1
		new_bytes = new_bytes / 1024
	}

	return fmt.Sprintf("%d%s", int(new_bytes), units[iteration])
}
