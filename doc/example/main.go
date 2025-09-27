package main

import (
	"context"

	"github.com/not-for-prod/observer/tracer"
	"github.com/not-for-prod/observer/tracer/prospan"
)

func bar(ctx context.Context) {
	ctx, span := prospan.Start(ctx)
	defer span.End()

	span.Logger().Info("bar")
}

func foo(ctx context.Context) {
	ctx, span := prospan.Start(ctx)
	defer span.End()

	span.Logger().Info("foo")

	bar(ctx)
}

func main() {
	ctx := context.Background()
	tp := tracer.NewProvider()

	// Инициализация трейсинга
	if err := tp.Start(ctx); err != nil {
		panic(err)
	}

	// graceful shutdown
	defer func() {
		if err := tp.Stop(ctx); err != nil {
			panic(err)
		}
	}()

	foo(context.Background())
}
