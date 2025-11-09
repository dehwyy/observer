package logger

import (
	"context"

	"go.uber.org/fx"
)

func NewLoggerFx(
	logger Logger,
) fx.Option {
	return fx.Invoke(
		func(lc fx.Lifecycle) {
			lc.Append(
				fx.Hook{
					OnStart: func(ctx context.Context) error {
						SetLogger(logger)
						return nil
					},
					OnStop: Stop,
				},
			)
		},
	)
}
