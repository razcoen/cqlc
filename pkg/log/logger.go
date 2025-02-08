package log

type Logger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
	With(args ...any) Logger
}

var _ Logger = (*nopLogger)(nil)

func NopLogger() Logger {
	return &nopLogger{}
}

type nopLogger struct{}

func (n nopLogger) Debug(msg string, args ...any) {}
func (n nopLogger) Info(msg string, args ...any)  {}
func (n nopLogger) Warn(msg string, args ...any)  {}
func (n nopLogger) Error(msg string, args ...any) {}
func (n nopLogger) With(args ...any) Logger       { return n }
