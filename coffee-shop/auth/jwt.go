package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// extracted to .env irl
// openssl rand -base64 32 (256 bits long)
var jwt_secret = []byte("oKqxelN8a8Cp5ibZ5K+9aBkWOdkc9EZXua8Dwi03bCc")

func CreateToken(username string) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		},
	)

	tokenString, err := token.SignedString(jwt_secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
