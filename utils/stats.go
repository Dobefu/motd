package utils

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
)

func GetOsVersion() string {
	os := GetOS()

	switch os {
	case "darwin":
		var stdout bytes.Buffer
		cmd := exec.Command("sw_vers")
		cmd.Stdout = &stdout
		err := cmd.Run()

		if err != nil {
			return "Unknown OS version"
		}

		regex, err := regexp.Compile("ProductVersion:\t+(.*)")

		if err != nil {
			return "Unknown OS version"
		}

		version_match := regex.FindStringSubmatch(stdout.String())

		if len(version_match) < 2 {
			return "MacOS"
		}

		return fmt.Sprintf("MacOS %s", version_match[1])
	default:
		return os
	}

}
