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

// Run resource availability tests and execute a given command
func Run(ctx context.Context, program Program, setters ...Option) ([]byte, error) {
	err := Test(ctx, program.Resources, setters...)

	if err != nil {
		return nil, err
	}

	cmd := exec.Command(program.Executable, program.Args...)

	return cmd.CombinedOutput()
}

// Run resource availability tests
func Test(ctx context.Context, resources []string, setters ...Option) error {
	opts := NewOptions(setters...)

	var buff bytes.Buffer
	output := testAllInternal(ctx, resources, *opts)

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

func testAllInternal(ctx context.Context, resources []string, opts Options) <-chan error {
	var wg sync.WaitGroup
	wg.Add(len(resources))

	output := make(chan error, len(resources))

	for _, resource := range resources {
		resource := resource

		go func() {
			defer wg.Done()

			output <- testInternal(ctx, resource, opts)
		}()
	}

	go func() {
		wg.Wait()
		close(output)
	}()

	return output
}

func testInternal(ctx context.Context, resource string, opts Options) error {
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
