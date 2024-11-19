package math

import (
	"testing"
)

func TestDuration(t *testing.T) {
	assertions := []struct {
		arg1   int
		arg2   int
		output int
	}{
		{1, 2, 2},
		{2, 1, 2},
		{1, 1, 1},
		{-2, 1, 1},
	}

	for _, assertion := range assertions {
		result := MaxInt(assertion.arg1, assertion.arg2)

		if result != assertion.output {
			t.Fatalf(
				"Expected (%d, %d) to match \"%d\", got \"%d\"",
				assertion.arg1,
				assertion.arg2,
				assertion.output,
				result,
			)
		}
	}
}
