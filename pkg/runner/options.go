package runner

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

func NewOptions() *Options {
	return &Options{
		interval:    time.Duration(5) * time.Second,
		maxInterval: time.Duration(60) * time.Second,
		attempts:    5,
	}
}

func WithInterval(interval uint64) Option {
	return func(opts *Options) {
		opts.interval = time.Duration(interval) * time.Second
	}
}

func WithMaxInterval(interval uint64) Option {
	return func(opts *Options) {
		opts.maxInterval = time.Duration(interval) * time.Second
	}
}

func WithAttempts(attempts uint64) Option {
	return func(opts *Options) {
		opts.attempts = attempts
	}
}
