package configurations

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() string {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	mongoDSN := os.Getenv("MONGO_DSN")
	if mongoDSN == "" {
		panic("MONGO_DSN is not set in .env file")
	}

	log.Println("MONGO_DSN: ", mongoDSN)

	return mongoDSN
}
