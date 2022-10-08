package utils

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// uniform error for env releted errors
func CheckEnv(envName string) string {
	value := os.Getenv(envName)
	if value == "" {
		log.Error("failed to start semi-task because required environment variable '%s' is not set", envName)
		os.Exit(-1)
	}
	return value
}
