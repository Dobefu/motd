package utils

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
)

func GetOsVersion() (string, error) {
	os := GetOS()

	switch os {
	case "darwin":
		var stdout bytes.Buffer
		cmd := exec.Command("sw_vers")
		cmd.Stdout = &stdout
		err := cmd.Run()

		if err != nil {
			return "", err
		}

		regex, err := regexp.Compile("ProductVersion:\t+(.*)")

		if err != nil {
			return "", err
		}

		version_match := regex.FindStringSubmatch(stdout.String())

		if len(version_match) < 2 {
			return "MacOS", nil
		}

		return fmt.Sprintf("MacOS %s", version_match[1]), nil
	default:
		return os, nil
	}

}
