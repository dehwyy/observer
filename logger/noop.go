package logger

import (
	"context"
	"log/slog"
)

type noop struct {
	logger *slog.Logger
}

func newNoop() Logger {
	return &noop{logger: slog.Default()}
}

func (n *noop) With(ctx context.Context, keysValues ...any) (context.Context, Logger) {
	ctx, kv := Upsert(ctx, keysValues...)
	return ctx, &noop{logger: slog.With(kv...)}
}

func (n *noop) Debug(msg string, keysValues ...any) { n.logger.Debug(msg, keysValues...) }
func (n *noop) Info(msg string, keysValues ...any)  { n.logger.Info(msg, keysValues...) }
func (n *noop) Warn(msg string, keysValues ...any)  { n.logger.Warn(msg, keysValues...) }
func (n *noop) Error(msg string, keysValues ...any) { n.logger.Error(msg, keysValues...) }
func (n *noop) Panic(msg string, keysValues ...any) { n.logger.Error(msg, keysValues...) }

func (n *noop) Sync() error {
	return nil
}
