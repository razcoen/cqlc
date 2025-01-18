package gocqlc

type Logger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
}

var _ Logger = (*NoopLogger)(nil)

type NoopLogger struct{}

func (n NoopLogger) Debug(msg string, args ...any) {}
func (n NoopLogger) Info(msg string, args ...any)  {}
func (n NoopLogger) Warn(msg string, args ...any)  {}
func (n NoopLogger) Error(msg string, args ...any) {}
