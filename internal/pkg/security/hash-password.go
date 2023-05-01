package secutiry

import "golang.org/x/crypto/bcrypt"

func GenerateHash(password string) (hash string, err error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return
	}
	hash = string(bytes)
	return
}

func CheckPasswordHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
