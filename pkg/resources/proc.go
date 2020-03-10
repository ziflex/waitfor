package resources

import (
	"context"
	"fmt"

	"github.com/mitchellh/go-ps"
)

type Process struct {
	name string
}

func NewProcess(name string) Resource {
	return &Process{name}
}

func (p *Process) Test(ctx context.Context) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}

	list, err := ps.Processes()

	if err != nil {
		return err
	}

	var found bool

	for _, proc := range list {
		if proc.Executable() == p.name {
			found = true
			break
		}
	}

	if found {
		return nil
	}

	return fmt.Errorf("process not found: %s", p.name)
}
