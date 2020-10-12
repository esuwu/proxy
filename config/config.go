package config

import (
	"os"
	"strconv"
)

type Config struct {
	BackendID      int64
	ServerEndpoint string
}

func New() Config {
	return Config{
		BackendID:      getIntEnvOrDefault("BACKEND_ID", 3),
		ServerEndpoint: getStringEnvOrDefault("SERVER_ENDPOINT", ":8081"),
	}
}

func getStringEnvOrDefault(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getIntEnvOrDefault(name string, defaultVal int64) int64 {
	valueStr := getStringEnvOrDefault(name, "")
	if value, err := strconv.ParseInt(valueStr, 10, 64); err == nil {
		return value
	}

	return defaultVal
}
