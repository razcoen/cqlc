package log

import "log/slog"

var _ Logger = (*SlogAdapter)(nil)

// SlogAdapter is an adapter that allows a slog.Logger to be used
// where a Logger interface is expected.
type SlogAdapter struct {
	SlogLogger *slog.Logger
}

// NewSlogAdapter creates a new SlogAdapter instance with the provided slog.Logger.
func NewSlogAdapter(logger *slog.Logger) *SlogAdapter {
	return &SlogAdapter{SlogLogger: logger}
}

// Debug logs a debug message using the underlying slog.Logger.
func (s *SlogAdapter) Debug(msg string, args ...any) { s.SlogLogger.Debug(msg, args...) }

// Error logs an error message using the underlying slog.Logger.
func (s *SlogAdapter) Error(msg string, args ...any) { s.SlogLogger.Error(msg, args...) }

// Info logs an informational message using the underlying slog.Logger.
func (s *SlogAdapter) Info(msg string, args ...any) { s.SlogLogger.Info(msg, args...) }

// Warn logs a warning message using the underlying slog.Logger.
func (s *SlogAdapter) Warn(msg string, args ...any) { s.SlogLogger.Warn(msg, args...) }

// With creates a new SlogAdapter with additional context using the underlying slog.Logger.
func (s *SlogAdapter) With(args ...any) Logger { return NewSlogAdapter(slog.With(args...)) }
