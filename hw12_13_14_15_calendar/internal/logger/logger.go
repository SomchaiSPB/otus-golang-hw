package logger

import (
	"fmt"
	"os"

	"github.com/SomchaiSPB/otus-golang-hw/hw12_13_14_15_calendar/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(conf config.LoggerConf) *zap.Logger {
	var level zapcore.Level
	pe := zap.NewProductionEncoderConfig()
	fileEncoder := zapcore.NewJSONEncoder(pe)
	consoleEncoder := zapcore.NewConsoleEncoder(pe)

	file, err := os.OpenFile(conf.File, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0o660)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	switch conf.Level {
	case "info":
		level = zap.InfoLevel
	case "debug":
		level = zap.DebugLevel
	default:
		level = zap.InfoLevel
	}

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, zapcore.AddSync(file), level),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), level),
	)

	logger := zap.New(core)
	defer logger.Sync()

	if err != nil {
		return nil
	}

	return logger
}
