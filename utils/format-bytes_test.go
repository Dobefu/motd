package utils

import (
	"testing"
)

func TestFormatBytes(t *testing.T) {
	assertions := []struct {
		input  int
		output string
	}{
		{1, "1 byte"},
		{10, "10 bytes"},
		{1000, "1000 bytes"},
		{1024, "1KiB"},
		{1025, "1KiB"},
		{2047, "1KiB"},
		{2048, "2KiB"},
		{1024 * 1024, "1MiB"},
		{1024 * 1024 * 1024, "1GiB"},
		{1024 * 1024 * 1024 * 1024, "1TiB"},
		{1024 * 1024 * 1024 * 1024 * 1024, "1PiB"},
		{1024 * 1024 * 1024 * 1024 * 1024 * 1024, "1EiB"},
	}

	for _, assertion := range assertions {
		result := FormatBytes(assertion.input)

		if result != assertion.output {
			t.Fatalf(
				"Expected %d to match \"%s\", got \"%s\"",
				int(assertion.input),
				assertion.output,
				result,
			)
		}
	}
}
