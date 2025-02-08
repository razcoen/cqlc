package log

type Logger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
	With(args ...any) Logger
}

var _ Logger = (*emptyLogger)(nil)

func EmptyLogger() Logger {
	return &emptyLogger{}
}

type emptyLogger struct{}

func (n emptyLogger) Debug(msg string, args ...any) {}
func (n emptyLogger) Info(msg string, args ...any)  {}
func (n emptyLogger) Warn(msg string, args ...any)  {}
func (n emptyLogger) Error(msg string, args ...any) {}
func (n emptyLogger) With(args ...any) Logger       { return n }
