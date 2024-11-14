package utils

import "runtime"

func GetOS() string {
	return runtime.GOOS
}
