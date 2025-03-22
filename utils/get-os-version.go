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
		cmd := exec.Command("/usr/bin/sw_vers")
		cmd.Stdout = &stdout
		err := cmd.Run()

		if err != nil {
			return "Unknown OS version"
		}

		regex, err := regexp.Compile("ProductVersion:\t+(.*)")

		if err != nil {
			return "Unknown OS version"
		}

		versionMatch := regex.FindStringSubmatch(stdout.String())

		if len(versionMatch) < 2 {
			return "MacOS"
		}

		return fmt.Sprintf("MacOS %s", versionMatch[1])
	default:
		return os
	}

}
