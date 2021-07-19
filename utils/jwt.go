package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func IDaasSignJWT(accessKey, secretKey string) (string, error) {
	claims := jwt.MapClaims{
		"exp":         time.Now().Add(time.Duration(10) * time.Minute).Unix(),
		"accessKeyId": accessKey,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}
