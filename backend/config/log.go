package config

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var logFile *os.File

func InitiateLog() {
	_, err := os.Stat("log")
	if os.IsNotExist(err) {
		os.Mkdir("log", os.ModePerm)
	}

	logFile, err = os.OpenFile("log/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to open log file")
	}

	log.Logger = zerolog.New(logFile).With().
		Timestamp().Logger()

	log.Info().Msg("Zerolog initiated successfully")
}

func CloseLog() {
	if logFile != nil {
		err := logFile.Close()
		if err != nil {
			log.Error().Err(err).Msg("Failed to close log file")
		}
	}
}
