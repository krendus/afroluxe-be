package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

type Env struct {
	DbName     string
	MongodbUri string
	PORT       string
}

func LoadEnv() Env {

	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error Loading .env file:: \n %v", err)
		}
	}
	loadedEnv := Env{
		DbName:     os.Getenv("DB_NAME"),
		MongodbUri: os.Getenv("MONGODB_URI"),
		PORT:       os.Getenv("PORT"),
	}

	return loadedEnv

}
