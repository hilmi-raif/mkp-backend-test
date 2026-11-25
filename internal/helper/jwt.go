package helper

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func SignToken(secret []byte, userID int64, expiresAt time.Time) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:   strconv.FormatInt(userID, 10),
		ExpiresAt: jwt.NewNumericDate(expiresAt),
	}).SignedString(secret)
}

func ValidToken(secret []byte, token string) bool {
	claims := &jwt.RegisteredClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (any, error) { return secret, nil },
		jwt.WithValidMethods([]string{"HS256"}), jwt.WithExpirationRequired())
	id, idErr := strconv.ParseInt(claims.Subject, 10, 64)
	return err == nil && idErr == nil && id > 0
}
