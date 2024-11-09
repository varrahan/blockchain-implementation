package utils

import (
	"os"
)

var EnvUtils map[string]string = map[string]string{
	"ADJUSTMENT_INTERVAL": os.Getenv("ADJUSTMENT_INTERVAL"),
	"CONTAINER_PORT":      os.Getenv("CONTAINER_PORT"),
	"INITIAL_DIFFICULTY":  os.Getenv("INITIAL_DIFFICULTY"),
	"HOST_PORT":           os.Getenv("HOST_PORT"),
	"SECURITY_CERT_PATH":  os.Getenv("SECURITY_CERT_PATH"),
	"SECURITY_KEY_PATH":   os.Getenv("SECURITY_KEY_PATH"),
	"TARGET_TIME":         os.Getenv("TARGET_TIME"),
	"TIMER_OFFSET":        os.Getenv("TIMER_OFFSET"),
}
