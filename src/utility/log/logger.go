package log

type Logger interface {
	Init()
	TrackTrace(string, ...string)
	TrackEvent(string, ...string)
	TrackError(error, ...string)
}
