package auth

import (
	"crypto/subtle"

	"golang.org/x/crypto/scrypt"
)

var salt = []byte("flooplewaffle")
var costFactor = 16384
var blockSize = 8
var parallelization = 1

func HashPassword(password string) (string, error) {
	key, err := scrypt.Key([]byte(password), salt, costFactor, blockSize, parallelization, 32)
	if err != nil {
		return "", err
	}
	return string(key), nil
}

func VerifyPassword(hashedPassword, password string) (bool, error) {
	hashedKey, err := scrypt.Key([]byte(password), salt, costFactor, blockSize, parallelization, 32)
	if err != nil {
		return false, err
	}
	//avoid timing attacks
	return subtle.ConstantTimeCompare(hashedKey, []byte(hashedPassword)) == 1, nil
}
