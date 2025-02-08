package log

import "github.com/charmbracelet/log"

var _ Logger = (*CharmbraceletAdapter)(nil)

// CharmbraceletAdapter is an adapter that allows a log.Logger to be used
// where a Logger interface is expected.
type CharmbraceletAdapter struct {
	CharmbraceletLogger *log.Logger
}

// NewCharmbraceletAdapter creates a new CharmbraceletAdapter instance with the provided log.Logger.
func NewCharmbraceletAdapter(logger *log.Logger) *CharmbraceletAdapter {
	return &CharmbraceletAdapter{CharmbraceletLogger: logger}
}

// Debug logs a debug message using the underlying log.Logger.
func (s *CharmbraceletAdapter) Debug(msg string, args ...any) {
	s.CharmbraceletLogger.Debug(msg, args...)
}

// Error logs an error message using the underlying log.Logger.
func (s *CharmbraceletAdapter) Error(msg string, args ...any) {
	s.CharmbraceletLogger.Error(msg, args...)
}

// Info logs an informational message using the underlying log.Logger.
func (s *CharmbraceletAdapter) Info(msg string, args ...any) {
	s.CharmbraceletLogger.Info(msg, args...)
}

// Warn logs a warning message using the underlying log.Logger.
func (s *CharmbraceletAdapter) Warn(msg string, args ...any) {
	s.CharmbraceletLogger.Warn(msg, args...)
}

// With creates a new CharmbraceletAdapter with additional context using the underlying log.Logger.
func (s *CharmbraceletAdapter) With(args ...any) Logger {
	return NewCharmbraceletAdapter(s.CharmbraceletLogger.With(args...))
}
