package utils

import (
	"motd/icons"
	"motd/structs"
)

func GetIcon(os string) (string, structs.RGB) {
	switch os {
	case "darwin":
		return icons.IconApple()
	default:
		return "", structs.RGB{255, 255, 255}
	}
}
