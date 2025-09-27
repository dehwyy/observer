package zap

import (
	"context"

	"github.com/unitewt/pkg/observability/logger"
	"go.uber.org/zap"
)

type Logger struct {
	logger *zap.SugaredLogger
}

func NewLogger() *Logger {
	cfg := zap.NewProductionConfig()
	base, _ := cfg.Build()
	return &Logger{
		logger: base.Sugar(),
	}
}

func (l *Logger) With(ctx context.Context, keysValues ...any) (context.Context, logger.Logger) {
	ctx, kv := logger.Upsert(ctx, keysValues...)
	zapLogger := l.logger.With(kv...)
	return ctx, &Logger{logger: zapLogger}
}

func (l *Logger) Debug(msg string, keysValues ...any) { l.logger.Debugw(msg, keysValues...) }
func (l *Logger) Info(msg string, keysValues ...any)  { l.logger.Infow(msg, keysValues...) }
func (l *Logger) Warn(msg string, keysValues ...any)  { l.logger.Warnw(msg, keysValues...) }
func (l *Logger) Error(msg string, keysValues ...any) { l.logger.Errorw(msg, keysValues...) }
func (l *Logger) Panic(msg string, keysValues ...any) { l.logger.Panicw(msg, keysValues...) }

func (l *Logger) Sync() error {
	_ = l.logger.Sync()
	return nil
}
