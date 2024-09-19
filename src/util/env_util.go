package util

import (
	"os"

	godotenv "github.com/joho/godotenv"
)

func EnvUtils() map[string]string {
	// load environment variables from .env
	godotenv.Load()

	// create map to store variables
	env_util := map[string]string{
		"PORT": os.Getenv("PORT_NUMBER"),
	}
	return env_util
}
