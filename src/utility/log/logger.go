package log

type Logger interface {
	Init()
	TrackTrace(string, ...string)
	TrackEvent(string, ...string)
	TrackError(error, ...string)
}

var initLoggerFunc map[string]func(name string) *Logger = map[string]func(name string) *Logger{
	"console": func(name string) *Logger {
		return NewConsoleLogger(name)
	},
}

func InitLogger() {

}
