package utils

import (
	"testing"
)

func TestDuration(t *testing.T) {
	assertions := []struct {
		input  float64
		output string
	}{
		{1, "1 second"},
		{10, "10 seconds"},
		{60, "1 minute"},
		{70, "1 minute 10 seconds"},
		{120, "2 minutes"},
		{121, "2 minutes 1 second"},
		{122, "2 minutes 2 seconds"},
		{3599, "59 minutes 59 seconds"},
		{3600, "1 hour"},
		{3601, "1 hour 1 second"},
		{3602, "1 hour 2 seconds"},
		{7199, "1 hour 59 minutes 59 seconds"},
		{7200, "2 hours"},
		{7201, "2 hours 1 second"},
		{7260, "2 hours 1 minute"},
		{7261, "2 hours 1 minute 1 second"},
		{7262, "2 hours 1 minute 2 seconds"},
		{86400, "1 day"},
		{86401, "1 day 1 second"},
		{86460, "1 day 1 minute"},
		{86461, "1 day 1 minute 1 second"},
		{86462, "1 day 1 minute 2 seconds"},
	}

	for _, assertion := range assertions {
		result := Duration(assertion.input)

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
