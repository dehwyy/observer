package logger

import "context"

//go:generate moq -out noop.go . Logger

type Logger interface {
	With(ctx context.Context, keysValues ...any) (context.Context, Logger)
	Debug(msg string, keysValues ...any)
	Info(msg string, keysValues ...any)
	Warn(msg string, keysValues ...any)
	Error(msg string, keysValues ...any)
	Panic(msg string, keysValues ...any)
	Sync() error
}
