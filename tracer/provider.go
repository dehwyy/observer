package tracer

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

type Provider struct {
	options  options
	exporter *otlptrace.Exporter
	provider *tracesdk.TracerProvider
}

func NewProvider(options ...Option) *Provider {
	opts := defaultOptions()

	for _, o := range options {
		o.apply(&opts)
	}

	return &Provider{
		options: opts,
	}
}

func (p *Provider) Start(ctx context.Context) error {
	var err error

	p.exporter, err = otlptracegrpc.New(
		ctx,
		otlptracegrpc.WithEndpoint(p.options.host),
		otlptracegrpc.WithInsecure(),
	)
	if err != nil {
		return err
	}

	p.provider = tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(p.exporter),
		tracesdk.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String(p.options.serviceName),
				semconv.ServiceVersionKey.String(p.options.serviceVersion),
			),
		),
	)

	otel.SetTracerProvider(p.provider)
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)

	return nil
}

func (p *Provider) Stop(ctx context.Context) error {
	if p.provider != nil {
		if err := p.provider.Shutdown(ctx); err != nil {
			return err
		}
	}

	if p.exporter != nil {
		return p.exporter.Shutdown(ctx)
	}

	return nil
}
