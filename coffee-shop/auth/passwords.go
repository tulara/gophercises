package auth

import (
	"crypto/subtle"

	"golang.org/x/crypto/scrypt"
)

// extracted to .env irl
// openssl rand -base64 20
var salt = []byte("Fk+1lifqB4erZDNe/A3TNEOB+HSuaH9/NnTdHn9lXmyWns")
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
