package log

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	log.Logger = newLogger()
}

func newLogger() zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	multiWriter := zerolog.MultiLevelWriter(zerolog.ConsoleWriter{Out: os.Stdout})
	return zerolog.New(multiWriter).With().Timestamp().Logger().Level(zerolog.DebugLevel)
}
