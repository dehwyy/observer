package tracer

import (
	"go.uber.org/fx"
)

func NewProviderFx(options ...Option) fx.Option {
	return fx.Invoke(
		func(LC fx.Lifecycle) *Provider {
			provider := NewProvider(options...)
			LC.Append(fx.Hook{
				OnStart: provider.Start,
				OnStop:  provider.Stop,
			})

			return provider
		},
	)
}
