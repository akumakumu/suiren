package utils

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret []byte
var jwtExpirationHours int

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	jwtSecret = []byte(os.Getenv("JWT_SECRET"))
	jwtExpirationHoursStr := os.Getenv("JWT_EXPIRATION")
	jwtExpirationHours, err = strconv.Atoi(jwtExpirationHoursStr)

	if err != nil {
		log.Printf("Invalid JWT Expiration: %s", err)
	}
}

func HashPassword(password string) (string, error) {
	godotenv.Load()

	hashCostStr := os.Getenv("HASH_COST")
	hashCost, err := strconv.Atoi(hashCostStr)

	if err != nil {
		return "", fmt.Errorf("invalid hash cost: %v", err)
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), hashCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateJWT(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * time.Duration(jwtExpirationHours)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
