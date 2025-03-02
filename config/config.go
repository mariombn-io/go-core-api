package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ProjectName    string
	ProjectVersion string
	Environment    string

	DatabaseHost     string
	DatabasePort     int
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string
	DatabaseSSLMode  string

	RedisHost     string
	RedisPort     int
	RedisPassword string
	RedisDB       int
}

var AppConfig Config

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found or error loading it")
	}

	AppConfig = Config{
		ProjectName:    os.Getenv("PROJECT_NAME"),
		ProjectVersion: os.Getenv("PROJECT_VERSION"),
		Environment:    os.Getenv("ENVIRONMENT"),

		DatabaseHost:     os.Getenv("DATABASE_HOST"),
		DatabaseUser:     os.Getenv("DATABASE_USER"),
		DatabasePassword: os.Getenv("DATABASE_PASSWORD"),
		DatabaseName:     os.Getenv("DATABASE_NAME"),
		DatabaseSSLMode:  os.Getenv("DATABASE_SSL_MODE"),

		RedisHost:     os.Getenv("REDIS_HOST"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
	}

	if portStr := os.Getenv("DATABASE_PORT"); portStr != "" {
		port, err := strconv.Atoi(portStr)
		if err != nil {
			log.Fatalf("Invalid DATABASE_PORT: %v", err)
		}
		AppConfig.DatabasePort = port
	}

	if redisPortStr := os.Getenv("REDIS_PORT"); redisPortStr != "" {
		port, err := strconv.Atoi(redisPortStr)
		if err != nil {
			log.Fatalf("Invalid REDIS_PORT: %v", err)
		}
		AppConfig.RedisPort = port
	}

	if redisDBStr := os.Getenv("REDIS_DB"); redisDBStr != "" {
		db, err := strconv.Atoi(redisDBStr)
		if err != nil {
			log.Fatalf("Invalid REDIS_DB: %v", err)
		}
		AppConfig.RedisDB = db
	}
}
