package utils

import "golang.org/x/crypto/bcrypt"

func HashPass(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func ComparePass(hashedPass, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(password))
	if err != nil {
		return false, err
	}

	return true, nil
}
