package logger

// Logger - common logger interface
type Logger interface {
	Debug(string, ...interface{})
	Info(string, ...interface{})
	Error(string, ...interface{})
	Fatal(string, ...interface{})
}
