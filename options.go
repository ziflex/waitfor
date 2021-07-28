package waitfor

import (
	"time"
)

type (
	Options struct {
		interval    time.Duration
		maxInterval time.Duration
		attempts    uint64
	}

	Option func(opts *Options)
)

// Create new options
func newOptions(setters []Option) *Options {
	opts := &Options{
		interval:    time.Duration(5) * time.Second,
		maxInterval: time.Duration(60) * time.Second,
		attempts:    5,
	}

	for _, setter := range setters {
		setter(opts)
	}

	return opts
}

// Set a custom test interval
func WithInterval(interval uint64) Option {
	return func(opts *Options) {
		opts.interval = time.Duration(interval) * time.Second
	}
}

// Set a custom maximum test interval
func WithMaxInterval(interval uint64) Option {
	return func(opts *Options) {
		opts.maxInterval = time.Duration(interval) * time.Second
	}
}

// Set a custom attempts count
func WithAttempts(attempts uint64) Option {
	return func(opts *Options) {
		opts.attempts = attempts
	}
}
