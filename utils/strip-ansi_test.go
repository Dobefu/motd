package utils

import (
	"fmt"
	"testing"

	"github.com/fatih/color"
)

func TestStripAnsi(t *testing.T) {
	color.NoColor = false

	assertions := []struct {
		input  string
		output string
	}{
		{"test", "test"},
		{color.RedString("test"), "test"},
		{fmt.Sprintf("%s %s", color.RedString("test"), color.RedString("test")), "test test"},
	}

	for _, assertion := range assertions {
		result := StripAnsi(assertion.input)

		if result != assertion.output {
			t.Fatalf(
				"Expected %s to match \"%s\", got \"%s\"",
				assertion.input,
				assertion.output,
				result,
			)
		}
	}
}
