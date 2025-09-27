package prospan

import (
	"context"
	"fmt"

	"github.com/unitewt/pkg/observability/logger"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type ProSpan struct {
	span   trace.Span
	logger logger.Logger
}

// Start starts the prospan.
func Start(ctx context.Context) (context.Context, ProSpan) {
	builder := New()
	builder.skipFrames++

	return builder.Start(ctx)
}

func (s *ProSpan) End(options ...trace.SpanEndOption) {
	s.span.End(options...)
}

func (s *ProSpan) Err(err error) error {
	s.span.SetStatus(codes.Error, fmt.Sprintf("%+v", err))
	s.span.RecordError(err)

	if s.logger != nil {
		s.logger.Error("prospan errored", "err", err.Error())
	}

	return err
}

func (s *ProSpan) Span() trace.Span {
	return s.span
}

func (s *ProSpan) TraceID() string {
	return s.span.SpanContext().TraceID().String()
}

func (s *ProSpan) Logger() logger.Logger {
	return s.logger
}
