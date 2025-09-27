package tracer

import "github.com/unitewt/pkg/git"

type options struct {
	host           string
	serviceName    string
	serviceVersion string
}

func defaultOptions() options {
	gitData := git.GetCommitInfo()

	return options{
		host:           "localhost:4317",
		serviceName:    gitData.Project,
		serviceVersion: gitData.String(),
	}
}

// Option overrides behavior of Connect.
type Option interface {
	apply(*options)
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options) {
	f(o)
}

func WithHost(url string) Option {
	return optionFunc(
		func(o *options) {
			o.host = url
		},
	)
}
