package port

type Logger interface {
	Trace(msg string, keyAndValues ...any)
	Debug(msg string, keyAndValues ...any)
	Info(msg string, keyAndValues ...any)
	Warning(msg string, keyAndValues ...any)
	Error(msg string, keyAndValues ...any)
	Fatal(msg string, keyAndValues ...any)
}
