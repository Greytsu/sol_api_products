package config

import (
	"github.com/rs/zerolog"
	"os"
)

func SetupLogger() {
	//Get log level from environment and set it to zerolog
	logLevel := os.Getenv("LOG_LEVEL")
	zerologLevel, err := zerolog.ParseLevel(logLevel)
	if err != nil {
		zerologLevel = zerolog.InfoLevel // Set a default log level
	}
	zerolog.SetGlobalLevel(zerologLevel)
}
