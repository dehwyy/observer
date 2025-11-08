package prospan

import (
	"context"

	"github.com/not-for-prod/observer/logger"
	"github.com/not-for-prod/observer/tracer/autoname"
	"go.opentelemetry.io/otel"
)

const initialSkipFrames = 2

// Builder дает добавть в спан какие то промежуточнык данные
type Builder struct {
	skipFrames int
	attributes map[string]any
}

func (b *Builder) Start(ctx context.Context) (context.Context, ProSpan) {
	ctx, span := otel.Tracer("").Start(ctx, autoname.GetRuntimeFunc(b.skipFrames))
	ctx, l := logger.Instance().With(ctx, "trace", span.SpanContext().TraceID().String())

	for key, val := range b.attributes {
		setAttr(span, key, val)
	}

	return ctx, ProSpan{
		span:   span,
		logger: l,
	}
}

func (b *Builder) WithAttribute(key string, value any) *Builder {
	b.attributes[key] = value
	return b
}

func WithAttribute(key string, value any) *Builder {
	return &Builder{
		attributes: map[string]any{key: value},
	}
}
