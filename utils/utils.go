package utils

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

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
