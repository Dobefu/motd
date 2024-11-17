package utils

import "os"

func GetTerminal() string {
	return os.Getenv("TERM")
}
