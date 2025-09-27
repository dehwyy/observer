package prospan

import (
	"context"

	"github.com/unitewt/pkg/observability/logger"
	"github.com/unitewt/pkg/observability/tracer/autoname"
	"go.opentelemetry.io/otel"
)

const initialSkipFrames = 2

// Builder дает добавть в спан какие то промежуточнык данные
type Builder struct {
	skipFrames int
}

func New() *Builder {
	return &Builder{
		skipFrames: initialSkipFrames,
	}
}

func (b Builder) Start(ctx context.Context) (context.Context, ProSpan) {
	ctx, span := otel.Tracer("").Start(ctx, autoname.GetRuntimeFunc(b.skipFrames))
	ctx, l := logger.Instance().With(ctx, "trace", span.SpanContext().TraceID().String())
	return ctx, ProSpan{
		span:   span,
		logger: l,
	}
}
