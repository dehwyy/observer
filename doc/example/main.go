package main

import (
	"context"

	"github.com/not-for-prod/observer/logger"
	"github.com/not-for-prod/observer/logger/zap"
	"github.com/not-for-prod/observer/tracer"
	"github.com/not-for-prod/observer/tracer/prospan"
	"golang.org/x/sync/errgroup"
)

func gar(ctx context.Context) {
	ctx, span := prospan.Start(ctx)
	defer span.End()

	span.Logger().Info("gar")
}

func bar(ctx context.Context) {
	ctx, span := prospan.Start(ctx)
	defer span.End()

	span.Logger().Info("bar")
}

func foo(ctx context.Context) error {
	ctx, span := prospan.WithAttribute("req", "test").Start(ctx)
	defer span.End()

	span.SetAttribute("custom-attr", 12345).
		SetAttribute("custom-attr", "custom-val").
		Logger().Info("foo")

	group, ctx := errgroup.WithContext(ctx)

	group.Go(
		func() error {
			bar(ctx)
			return nil
		},
	)
	group.Go(
		func() error {
			gar(ctx)
			return nil
		},
	)

	return group.Wait()
}

func main() {
	ctx := context.Background()
	tp := tracer.NewProvider()

	logger.SetLogger(zap.NewLogger())
	logger.Instance().Info("starting application")

	defer func() {
		if err := logger.Stop(ctx); err != nil {
			panic(err)
		}
	}()

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
