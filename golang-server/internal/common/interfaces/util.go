package interfaces

import "log"

type LogLevel string

const (
	INFO  LogLevel = "INFO"
	DEBUG LogLevel = "DEBUG"
	WARN  LogLevel = "WARN"
	ERROR LogLevel = "ERROR"
)

type LogData map[string]any

type LogUtil interface {
	Info(message string, data LogData)
	Debug(message string, data LogData)
	Warn(message string, data LogData)
	Error(message string, data LogData)
}

type Util struct {
	Log    LogUtil
	Logger *log.Logger
}
