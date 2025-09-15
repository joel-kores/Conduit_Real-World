package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

type Logger interface {
	Info(msg string, fields map[string]interface{})
	Error(msg string, fields map[string]interface{})
	Warn(msg string, fields map[string]interface{})
	Debug(msg string, field map[string]interface{})
	With(fields map[string]interface{}) Logger
}

type ZeroLogger struct {
	logger zerolog.Logger
}

func NewZeroLogger(level string) *ZeroLogger {
	zerologLevel, err := zerolog.ParseLevel(level)
	if err != nil {
		zerologLevel = zerolog.InfoLevel
	}
	zerolog.TimeFieldFormat = time.RFC3339Nano
	logger := zerolog.New(os.Stdout).
		Level(zerologLevel).
		With().
		Timestamp().
		Logger()

	return &ZeroLogger{logger: logger}
}

func (l *ZeroLogger) Info(msg string, fields map[string]interface{}) {
	l.logger.Info().Fields(fields).Msg(msg)
}

func (l *ZeroLogger) Error(msg string, fields map[string]interface{}) {
	l.logger.Error().Fields(fields).Msg(msg)
}

func (l *ZeroLogger) Warn(msg string, fields map[string]interface{}) {
	l.logger.Warn().Fields(fields).Msg(msg)
}

func (l *ZeroLogger) Debug(msg string, fields map[string]interface{}) {
	l.logger.Debug().Fields(fields).Msg(msg)
}

func (l *ZeroLogger) With(fields map[string]interface{}) Logger {
	return &ZeroLogger{logger: l.logger.With().Fields(fields).Logger()}
}
