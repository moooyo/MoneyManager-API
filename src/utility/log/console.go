package log

import (
	"fmt"
	"os"
)

type ConsoleLogger struct {
	EventBuffer *LoggerBuffer[*EventInfo]
	TraceBuffer *LoggerBuffer[*TraceInfo]
}

// implement ConsoleLogger to Logger
func ConsoleLogEvent(eventInfos []*EventInfo) {
	for _, event := range eventInfos {
		_, _ = fmt.Fprintln(os.Stdout, event.ToString())
	}
}

func ConsoleLogError(errorInfo *ErrorInfo) {
	_, _ = fmt.Fprintln(os.Stderr, errorInfo.ToString())
}

func ConsoleLogTrace(traceInfos []*TraceInfo) {
	for _, traceInfo := range traceInfos {
		_, _ = fmt.Fprintln(os.Stdout, traceInfo.ToString())
	}
}

func (logger *ConsoleLogger) TrackEvent(eventName string, args ...string) {
	event := &EventInfo{
		Name:             eventName,
		CustomDimensions: map[string]string{},
	}
	for i := 0; i < len(args)/2; i++ {
		event.CustomDimensions[args[i]] = args[i+1]
	}

	logger.EventBuffer.Write(event)
}

func (logger *ConsoleLogger) TrackError(err error, args ...string) {
	errInfo := &ErrorInfo{
		Error:            err,
		CustomDimensions: map[string]string{},
	}
	for i := 0; i < len(args)/2; i++ {
		errInfo.CustomDimensions[args[i]] = args[i+1]
	}

	ConsoleLogError(errInfo)
}

func (logger *ConsoleLogger) TrackTrace(message string, args ...string) {
	trace := &TraceInfo{
		Message:          message,
		CustomDimensions: map[string]string{},
	}
	for i := 0; i < len(args)/2; i++ {
		trace.CustomDimensions[args[i]] = args[i+1]
	}

	logger.TraceBuffer.Write(trace)
}

func NewConsoleLogger() Logger {
	return &ConsoleLogger{
		EventBuffer: MakeBuffer(100, 1, ConsoleLogEvent),
		TraceBuffer: MakeBuffer(100, 1, ConsoleLogTrace),
	}
}
