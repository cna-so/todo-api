package helpers

import "golang.org/x/crypto/bcrypt"

func GenerateHashPassword(pass string) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(pass), 8)
	if err != nil {
		return "", err
	}
	return string(password), err
}
