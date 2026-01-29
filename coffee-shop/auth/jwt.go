package auth

import (
	"errors"
	"fmt"
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

// VerifyJWTToken returns username where the token is valid or an error otherwise.
func VerifyJWTToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		// avoid ALgorithm Confusion attacks.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return jwt_secret, nil
	})

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", errors.New("invalid token")
	}

	var username string
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if u, ok := claims["username"].(string); ok {
			username = u
		}
	}

	if username == "" {
		return "", errors.New("invalid token claims")
	}

	return username, nil
}
