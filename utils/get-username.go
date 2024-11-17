package utils

import (
	"os/user"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func GetUsername() string {
	current_user, err := user.Current()

	if err != nil {
		return "Error"
	}

	return cases.Title(language.English, cases.Compact).String(current_user.Username)
}
