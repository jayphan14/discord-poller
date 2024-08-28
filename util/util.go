package util

import (
	"fmt"
	"os"
)

// GetEnv retrieves the environment variable, or returns the default value if it's not set
func GetEnv(key, defaultValue string, required bool) (string, error) {
	if value, exists := os.LookupEnv(key); exists {
		return value, nil
	}

	if required {
		return "", fmt.Errorf("Can't find %v in env var", key)
	}
	return defaultValue, nil
}
