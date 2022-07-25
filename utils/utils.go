package utils

import (
	"time"

	"github.com/afroluxe/afroluxe-be/config"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	Id string
	jwt.StandardClaims
}

func HashPassword(password string) (string, error) {
	byte, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(byte), err
}
func CheckPasswordHash(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateNewToken(id string, expTime time.Time) (string, error) {
	claims := &Claims{
		Id: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedStr, err := token.SignedString([]byte(config.LoadEnv().JwtKey))

	if err != nil {
		return "", err
	}
	return signedStr, nil
}

func VerifyToken(tokenStr string) (bool, string) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.LoadEnv().JwtKey), nil
	})

	if err != nil {
		return false, ""
	}
	if !token.Valid {
		return false, ""
	}
	return true, claims.Id
}
