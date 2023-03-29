package helpers

import (
	"log"

	"github.com/joho/godotenv"
)

func LogIfError(err error) {
	if err != nil {
		log.Println("[ERROR] ", err)
		panic(err)
	}
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("[ERROR] Error loading .env file")
		panic(err)
	}
}
