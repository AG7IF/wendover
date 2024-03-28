package logging

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}

func InitLogging(loglevel string, console bool) {
	if console {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
	lvl, err := zerolog.ParseLevel(loglevel)
	if err != nil {
		log.Error().Stack().Err(err).Str("loglevel", loglevel).Msg("failed to parse loglevel, defaulting to INFO")
		lvl = zerolog.InfoLevel
	}

	zerolog.SetGlobalLevel(lvl)
}
