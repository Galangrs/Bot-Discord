package config

import (
	"os"
)

// Define constants for environment variable names
const (
	DISCORD_TOKEN      = "DISCORD_TOKEN"
	HOST_DATABASE      = "HOST_DATABASE"
	USER_DATABASE      = "USER_DATABASE"
	PASSWORD_DATABASE  = "PASSWORD_DATABASE"
	NAME_DATABASE      = "NAME_DATABASE"
)

type Config struct {
	Host     string
	User     string
	Password string
	DBName   string
}

func Token() string {
	return os.Getenv(DISCORD_TOKEN)
}

func GetConfig() Config {
	return Config{
		Host:     os.Getenv(HOST_DATABASE),
		User:     os.Getenv(USER_DATABASE),
		Password: os.Getenv(PASSWORD_DATABASE),
		DBName:   os.Getenv(NAME_DATABASE),
	}
}
