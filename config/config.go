package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

type Env struct {
	DbName       string
	MongodbUri   string
	JwtKey       string
	PORT         string
	MailHost     string
	MailPort     string
	MailPassword string
	RedisUrl     string
}

func LoadEnv() Env {

	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error Loading .env file:: \n %v", err)
		}
	}
	loadedEnv := Env{
		DbName:       os.Getenv("DB_NAME"),
		MongodbUri:   os.Getenv("MONGODB_URI"),
		JwtKey:       os.Getenv("JWT_KEY"),
		PORT:         os.Getenv("PORT"),
		MailHost:     os.Getenv("MAIL_HOST"),
		MailPort:     os.Getenv("MAIL_PORT"),
		MailPassword: os.Getenv("MAIL_PASSWORD"),
		RedisUrl:     os.Getenv("REDIS_URL"),
	}

	return loadedEnv

}
