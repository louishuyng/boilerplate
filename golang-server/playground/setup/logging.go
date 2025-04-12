package playground_setup

import (
	"log"
	"log/slog"
	"os"
	"rz-server/internal/common/interfaces"
)

var _ interfaces.LogUtil = (*Log)(nil)

type Log struct {
	logger      *slog.Logger
	ErrorLogger *log.Logger
}

func NewLog() *Log {
	handler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(handler)
	errorLogger := slog.NewLogLogger(handler, slog.LevelError)

	slog.SetDefault(logger)

	return &Log{
		logger:      logger,
		ErrorLogger: errorLogger,
	}
}

func (l *Log) Info(message string, data ...interfaces.LogData) {
	l.logger.Info(message, "data", data)
}

func (l *Log) Debug(message string, data ...interfaces.LogData) {
	l.logger.Debug(message, "data", data)
}

func (l *Log) Warn(message string, data ...interfaces.LogData) {
	l.logger.Warn(message, "data", data)
}

func (l *Log) Error(message string, data ...interfaces.LogData) {
	l.logger.Error(message, "data", data)
}
