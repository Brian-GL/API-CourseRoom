package middleware

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func ValidateToken(token *string) bool {

	// Cargar archivo .env
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading token info")
		return false
	}

	//Cargar variable token desde env
	envToken := os.Getenv("SECRET_TOKEN")

	return envToken == *token
}
