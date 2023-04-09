package log

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

func TrackTrace(ctx *context.Context, message string, args ...string) {
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

func TrackEvent(context *fiber.Ctx, eventName string, args ...string) {
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

func TrackError(context *fiber.Ctx, err error, args ...string) {
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
