package log

import "github.com/gin-gonic/gin"

func TrackTrace(context *gin.Context, message string, args ...string) {
	val, exist := context.Get("logger")
	if !exist {
		panic("logger not exist")
	}
	logger, ok := val.(Logger)
	if !ok {
		panic("logger type error")
	}

	logger.TrackTrace(message, args...)
}

func TrackEvent(context *gin.Context, eventName string, args ...string) {
	val, exist := context.Get("logger")
	if !exist {
		panic("logger not exist")
	}
	logger, ok := val.(Logger)
	if !ok {
		panic("logger type error")
	}

	logger.TrackEvent(eventName, args...)
}

func TrackError(context *gin.Context, err error, args ...string) {
	val, exist := context.Get("logger")
	if !exist {
		panic("logger not exist")
	}
	logger, ok := val.(Logger)
	if !ok {
		panic("logger type error")
	}

	logger.TrackError(err, args...)
}
