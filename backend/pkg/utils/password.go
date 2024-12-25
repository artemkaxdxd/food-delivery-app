package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePasswords(hashPwd, pwd string) bool {
	password := []byte(pwd)
	hash := []byte(hashPwd)
	err := bcrypt.CompareHashAndPassword(hash, password)
	return err == nil
}
