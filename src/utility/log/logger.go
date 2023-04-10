package log

import (
	"MoneyManager/src/utility/utils"
	"sync"
)

type Logger interface {
	TrackTrace(string, ...string)
	TrackEvent(string, ...string)
	TrackError(error, ...string)
}

var initLoggerFunc map[string]func() Logger = map[string]func() Logger{
	"console": NewConsoleLogger,
}

var logger Logger
var loggerDoOnce sync.Once

func InitLogger() {
	cfg := utils.GetConfig()
	if f, ok := initLoggerFunc[cfg.App.LogType]; ok {
		logger = f()
		return
	}

	panic("Invalid log type")
}

func GetLogger() Logger {
	loggerDoOnce.Do(func() {
		InitLogger()
	})

	return logger
}
