package utils

import (
	"testing"

	"github.com/fatih/color"
)

func TestTrimAnsiString(t *testing.T) {
	assertions := []struct {
		arg1   string
		arg2   int
		output string
	}{
		{"test", 2, "te"},
		{"test", 4, "test"},
		{"test", 5, "test"},
		{"test", 0, ""},
		{color.RedString("test"), 4, color.RedString("test")},
		{color.RedString("test"), 5, color.RedString("test")},
		{color.RedString("test"), 2, color.RedString("te")},
	}

	for _, assertion := range assertions {
		result := TrimAnsiString(assertion.arg1, assertion.arg2)

		if result != assertion.output {
			t.Fatalf(
				"Expected (%s, %d) to match \"%s\", got \"%s\"",
				assertion.arg1,
				assertion.arg2,
				assertion.output,
				result,
			)
		}
	}
}
