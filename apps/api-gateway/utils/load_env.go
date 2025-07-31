package utils

import (
	// "fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system env")
	}
}

func Env(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		// fmt.Println(value);
		// fmt.Println(exists);
		// http://localhost:4000
		// true
		return value
	}
	return fallback
}
