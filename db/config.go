package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config func to get env value
func GetKeyConfig(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(key)
}

// struct all env
type (
	Config struct {
		Server ServerConfig
		DB     DBConfig
		//JWTToken JWTConfig
	}
	ServerConfig struct {
		Port string
	}
	DBConfig struct {
		Protocol string
		Host     string
		Port     string
		Username string
		Password string
		Name     string
	}
)

// Get all env
func LoadConfig() (config *Config, err error) {
	err = godotenv.Load(".env")
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	config = &Config{
		Server: ServerConfig{
			Port: os.Getenv("SERVER_PORT"),
		},
		DB: DBConfig{
			Protocol: os.Getenv("DB_PROTOCOL"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
		},
	}

	return config, nil
}
