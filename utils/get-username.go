package utils

import (
	"os/user"
)

func GetUsername() string {
	current_user, err := user.Current()

	if err != nil {
		return "Error"
	}

	return current_user.Username
}
