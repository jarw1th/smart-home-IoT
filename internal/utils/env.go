package utils

import "os"

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func GetPort() string {
	return getEnv("APP_PORT", "8080")
}
