package dmc

import (
	"io"

	"github.com/rs/zerolog"
)

type Logger struct {
	log zerolog.Logger
}

func logWriter() zerolog.ConsoleWriter {
	var output io.Writer = io.Discard

	return zerolog.ConsoleWriter{Out: output, TimeFormat: ISO8601}
}

func NewLogger() *Logger {
	return &Logger{
		log: zerolog.New(logWriter()),
	}
}

func (logger *Logger) LogError(err error, msg string) error {
	if err != nil {
		logger.log.Error().Err(err).Msg(msg)
	}

	return err
}
