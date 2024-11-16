package utils

import (
	"motd/icons"
)

func GetIcon(os string) string {
	switch os {
	case "darwin":
		return icons.IconApple()
	default:
		return ""
	}
}
