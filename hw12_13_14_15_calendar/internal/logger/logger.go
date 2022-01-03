package logger

import (
	"go.uber.org/zap"
)

type Logger struct { // TODO
	logger *zap.Logger
	level  string
}

func New(level string) *Logger {
	logger, err := zap.NewProduction()
	defer logger.Sync()

	if err != nil {
		return nil
	}

	return &Logger{
		logger: logger,
		level:  level,
	}
}

func (l Logger) Info(msg string) {
	l.logger.Info(msg)
}

func (l Logger) Error(msg string) {
	l.logger.Error(msg)
}

func (l *Logger) Fatal(msg string) {
	l.logger.Fatal(msg)
}
