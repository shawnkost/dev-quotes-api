package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	if os.Getenv("ENVIRONMENT") != "production" {
		log.Logger = log.Output(zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		})
	}

	level := zerolog.InfoLevel
	if os.Getenv("LOG_LEVEL") != "" {
		level, _ = zerolog.ParseLevel(os.Getenv("LOG_LEVEL"))
	}
	zerolog.SetGlobalLevel(level)

	zerolog.TimeFieldFormat = time.RFC3339
}

func Logger() zerolog.Logger {
	return log.With().Str("service", "dev-quotes-api").Logger()
}
