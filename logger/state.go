package logger

import (
	"context"
	"sync"
	"sync/atomic"
)

type loggerHandler struct {
	logger Logger
}

var (
	delegateLoggerOnce sync.Once
)

var (
	globalLogger = defaultLogger()
)

func defaultLogger() *atomic.Value {
	v := &atomic.Value{}
	v.Store(&loggerHandler{logger: newNoop()})
	return v
}

func SetLogger(logger Logger) {
	delegateLoggerOnce.Do(
		func() {
			globalLogger.Store(&loggerHandler{logger: logger})
		},
	)
}

func Instance() Logger {
	return globalLogger.Load().(*loggerHandler).logger
}

func Stop(_ context.Context) error {
	return Instance().Sync()
}
