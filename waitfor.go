package waitfor

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"sync"

	"github.com/cenkalti/backoff"
)

type (
	Program struct {
		Executable string
		Args       []string
		Resources  []string
	}

	Runner struct {
		registry *Registry
	}
)

func New(resources ...ResourceConfig) *Runner {
	r := new(Runner)
	r.registry = newRegistry(resources)

	return r
}

func (r *Runner) Resources() *Registry {
	return r.registry
}

// Run runs resource availability tests and execute a given command
func (r *Runner) Run(ctx context.Context, program Program, setters ...Option) ([]byte, error) {
	err := r.Test(ctx, program.Resources, setters...)

	if err != nil {
		return nil, err
	}

	cmd := exec.Command(program.Executable, program.Args...)

	return cmd.CombinedOutput()
}

// Test tests resource availability
func (r *Runner) Test(ctx context.Context, resources []string, setters ...Option) error {
	opts := newOptions(setters)

	var buff bytes.Buffer
	output := r.testAllInternal(ctx, resources, *opts)

	for err := range output {
		if err != nil {
			buff.WriteString(err.Error() + ";")
		}
	}

	if buff.Len() != 0 {
		return fmt.Errorf("failed to wait for resources: %s", buff.String())
	}

	return nil
}

func (r *Runner) testAllInternal(ctx context.Context, resources []string, opts Options) <-chan error {
	var wg sync.WaitGroup
	wg.Add(len(resources))

	output := make(chan error, len(resources))

	for _, resource := range resources {
		resource := resource

		go func() {
			defer wg.Done()

			output <- r.testInternal(ctx, resource, opts)
		}()
	}

	go func() {
		wg.Wait()
		close(output)
	}()

	return output
}

func (r *Runner) testInternal(ctx context.Context, resource string, opts Options) error {
	rsc, err := r.registry.Resolve(resource)

	if err != nil {
		return err
	}

	b := backoff.NewExponentialBackOff()
	b.InitialInterval = opts.interval
	b.MaxInterval = opts.maxInterval

	return backoff.Retry(func() error {
		return rsc.Test(ctx)
	}, backoff.WithContext(backoff.WithMaxRetries(b, opts.attempts), ctx))
}
