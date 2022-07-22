package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	DB_NAME     string
	MONGODB_URI string
	PORT        string
}

func LoadEnv() Env {

	if os.Getenv("APP_ENV") != "production" {
		_ = godotenv.Load()

		loadedEnv := Env{
			DB_NAME:     os.Getenv("DB_NAME"),
			MONGODB_URI: os.Getenv("MONGODB_URI"),
			PORT:        os.Getenv("PORT"),
		}

		return loadedEnv
	} else {
		loadedEnv := Env{
			DB_NAME:     os.Getenv("DB_NAME"),
			MONGODB_URI: os.Getenv("MONGODB_URI"),
			PORT:        os.Getenv("PORT"),
		}

		return loadedEnv
	}

}
