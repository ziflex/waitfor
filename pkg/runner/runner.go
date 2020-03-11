package runner

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"sync"

	"github.com/cenkalti/backoff"

	"github.com/ziflex/waitfor/pkg/resources"
)

type Program struct {
	Executable string
	Args       []string
	Resources  []string
}

func Run(ctx context.Context, program Program, setters ...Option) ([]byte, error) {
	opts := NewOptions()

	for _, setter := range setters {
		setter(opts)
	}

	var buff bytes.Buffer
	output := runAll(ctx, program.Resources, *opts)

	for err := range output {
		if err != nil {
			buff.WriteString(err.Error() + ";")
		}
	}

	if buff.Len() != 0 {
		return nil, fmt.Errorf("failed to wait for resources: %s", buff.String())
	}

	cmd := exec.Command(program.Executable, program.Args...)

	return cmd.CombinedOutput()
}

func runAll(ctx context.Context, resources []string, opts Options) <-chan error {
	var wg sync.WaitGroup
	wg.Add(len(resources))

	output := make(chan error, len(resources))

	for _, resource := range resources {
		resource := resource

		go func() {
			defer wg.Done()

			output <- run(ctx, resource, opts)
		}()
	}

	go func() {
		wg.Wait()
		close(output)
	}()

	return output
}

func run(ctx context.Context, resource string, opts Options) error {
	r, err := resources.New(resource)

	if err != nil {
		return err
	}

	b := backoff.NewExponentialBackOff()
	b.InitialInterval = opts.interval
	b.MaxInterval = opts.maxInterval

	return backoff.Retry(func() error {
		return r.Test(ctx)
	}, backoff.WithContext(backoff.WithMaxRetries(b, opts.attempts), ctx))
}
