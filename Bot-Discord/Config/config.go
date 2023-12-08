package config

import (
	"os"
)

// Define constants for environment variable names
const (
	DISCORD_TOKEN      	= "DISCORD_TOKEN"
	DISCORD_ID_OWNER	= "DISCORD_ID_OWNER"
	URL_FETCH			= "URL_FETCH"
)

func Token() string {
	return os.Getenv(DISCORD_TOKEN)
}

func URL() string {
	return os.Getenv(URL_FETCH)
}

func Owner() string {
	return os.Getenv(DISCORD_ID_OWNER)
}
