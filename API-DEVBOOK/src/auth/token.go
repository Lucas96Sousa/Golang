package auth

import (
	"api/src/config"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

// CreateToken - return a token assigned with user perm
func CreateToken(userID uint64) (string, error) {
	perm := jwt.MapClaims{}
	perm["authorized"] = true
	perm["exp"] = time.Now().Add(time.Hour * 2).Unix()
	perm["userID"] = userID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, perm)
	return token.SignedString([]byte(config.SecretKey)) //secret

}
