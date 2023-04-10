package log

import (
	"context"
)

func TrackTrace(ctx context.Context, message string, args ...string) {
	logger := GetLogger()

	logger.TrackTrace(message, args...)
}

func TrackEvent(ctx context.Context, eventName string, args ...string) {
	logger := GetLogger()

	logger.TrackEvent(eventName, args...)
}

func TrackError(ctx context.Context, err error, args ...string) {
	logger := GetLogger()

	logger.TrackError(err, args...)
}
