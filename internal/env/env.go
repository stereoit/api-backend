package env

import "os"

// Get read ENV variable with fallback to default
func Get(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
