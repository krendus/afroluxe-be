package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	DB_NAME     string
	MONGODB_URI string
	PORT        string
}

func LoadEnv() Env {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error Loading .env file:: \n %v", err)
	}
	loadedEnv := Env{
		DB_NAME:     os.Getenv("DB_NAME"),
		MONGODB_URI: os.Getenv("MONGODB_URI"),
		PORT:        os.Getenv("PORT"),
	}
	return loadedEnv
}
