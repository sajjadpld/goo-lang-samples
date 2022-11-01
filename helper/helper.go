package helper

import "strings"

func ValidateEmain(email string) bool {
	return strings.Contains(email, "@")
}
