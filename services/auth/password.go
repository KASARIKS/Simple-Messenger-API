package auth

import (
	"crypto/sha256"
	"fmt"
)

var IncorrectPassword error = fmt.Errorf("incorrect password")

func HashPassword(password string) (string, error) {
	hasher := sha256.New()
	_, err := hasher.Write([]byte(password))
	if err != nil {
		return "", err
	}

	hash := hasher.Sum(nil)

	return string(hash), nil
}

func ComparePasswords(currentHashedPassword, gottenPassword string) error {
	hashed, err := HashPassword(gottenPassword)
	if err != nil {
		return err
	}

	if currentHashedPassword != hashed {
		return IncorrectPassword
	}

	return nil
}
