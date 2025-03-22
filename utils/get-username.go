package utils

import (
	"os/user"
)

func GetUsername() string {
	currentUser, err := user.Current()

	if err != nil {
		return "Error"
	}

	return currentUser.Username
}
