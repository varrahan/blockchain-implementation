package utils

import (
	"os"
	"strconv"

	godotenv "github.com/joho/godotenv"
)

func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

func IntToString(i int) string {
	return strconv.Itoa(i)
}

func StringToFloat(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

func StringToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func EnvUtils() map[string]string {
	// load environment variables from .env
	godotenv.Load()
	// create map to store variables
	env_util := map[string]string{
		"ADJUSTMENT_INTERVAL": os.Getenv("ADJUSTMENT_INTERVAL"),
		"CONTAINER_PORT":      os.Getenv("CONTAINER_PORT"),
		"INITIAL_DIFFICULTY":  os.Getenv("INITIAL_DIFFICULTY"),
		"HOST_PORT":           os.Getenv("HOST_PORT"),
		"TARGET_TIME":         os.Getenv("TARGET_TIME"),
		"TIMER_OFFSET":        os.Getenv("TIMER_OFFSET"),
	}
	return env_util
}
