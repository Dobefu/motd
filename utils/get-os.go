package utils

import "runtime"

func GetOS() string {
	os := runtime.GOOS

	if os != "linux" {
		return os
	}

	return "TODO"
}
