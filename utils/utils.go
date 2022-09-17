package utils

import (
	"log"
	"math/rand"
	"os"
	"strconv"
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

func CreateNewToken(id string) (string, error) {
	duration, err := strconv.Atoi(loadedEnv.JwtDuration)
	if err != nil {
		logger := log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Println("Invalid JWT Duration")
	}
	claims := &Claims{
		Id: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(duration)).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedStr, err := token.SignedString([]byte(config.LoadEnv().JwtKey))

	if err != nil {
		return "", err
	}
	return signedStr, nil
}

func VerifyToken(tokenStr string) (bool, *Claims) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.LoadEnv().JwtKey), nil
	})
	if err != nil {
		return false, claims
	}
	if !token.Valid {
		return false, claims
	}

	return true, claims
}

func GenerateRandomOtp(size uint8) string {
	newRand := rand.New(rand.NewSource(time.Now().Unix()))
	table := [10]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, size)
	for i := uint8(0); i < size; i++ {
		b[i] = table[newRand.Intn(len(table))]
	}
	return string(b)
}

func InfoLogger(message interface{}) {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println(message)
}

func WarningLogger(message interface{}) {
	logger := log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println(message)
}

func ErrorLogger(message interface{}) {
	logger := log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println(message)
}
